package elasticsearch

import (
	"bytes"
	"context"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"diffme.dev/diffme-api/internal/shared/encoders"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/wI2L/jsondiff"
	"log"
	"strings"
	"time"
)

var (
	changeIndex = "changes"
)

type ChangeSearchRepository struct {
	client *elasticsearch.Client
}

type Diff jsondiff.Operation

type SearchChangeModel struct {
	Id          string                 `json:"id" mapstructure:"id"`
	ChangeSetId string                 `json:"change_set_id" mapstructure:"change_set_id"`
	SnapshotId  string                 `json:"snapshot_id" mapstructure:"snapshot_id"`
	ReferenceId string                 `json:"reference_id" mapstructure:"reference_id"`
	Editor      string                 `json:"editor" mapstructure:"editor"`
	Metadata    map[string]interface{} `json:"metadata" mapstructure:"metadata"`
	Diff        Diff                   `json:"diff" mapstructure:"diff"`
	UpdatedAt   time.Time              `json:"updated_at" mapstructure:"updated_at"`
	CreatedAt   time.Time              `json:"created_at" mapstructure:"created_at"`
}

func NewElasticSearchChangeRepo(client *elasticsearch.Client) domain.SearchChangeRepository {
	return &ChangeSearchRepository{client: client}
}

func (m *ChangeSearchRepository) toDomain(doc SearchChangeModel) domain.SearchChange {
	return domain.SearchChange{
		Id:          doc.Id,
		ChangeSetId: doc.ChangeSetId,
		ReferenceId: doc.ReferenceId,
		SnapshotId:  doc.SnapshotId,
		Editor:      doc.Editor,
		Metadata:    doc.Metadata,
		Diff:        domain.Diff(doc.Diff),
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *ChangeSearchRepository) toPersistence(change domain.SearchChange) SearchChangeModel {
	return SearchChangeModel{
		Id:          change.Id,
		ChangeSetId: change.ChangeSetId,
		ReferenceId: change.ReferenceId,
		SnapshotId:  change.SnapshotId,
		Editor:      change.Editor,
		Metadata:    change.Metadata,
		Diff:        Diff(change.Diff),
		UpdatedAt:   change.UpdatedAt,
		CreatedAt:   change.CreatedAt,
	}
}

func (m *ChangeSearchRepository) Query(match domain.SearchRequest) ([]domain.SearchChange, error) {

	var (
		results map[string]interface{}
		buf     bytes.Buffer
	)

	var must []interface{}

	if match.Editor != nil {
		editorMatch := map[string]interface{}{
			"match": map[string]interface{}{
				"editor": match.Editor,
			},
		}
		must = append(must, editorMatch)
	}

	if match.Field != nil {
		fieldMatch := map[string]interface{}{
			"match": map[string]interface{}{
				"diff.path": match.Field,
			},
		}
		must = append(must, fieldMatch)
	}

	if match.Value != nil {
		valueMatch := map[string]interface{}{
			"match": map[string]interface{}{
				"diff.value": match.Value,
			},
		}
		must = append(must, valueMatch)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}

	//str, _ := json.MarshalIndent(query, "", "  ")
	//fmt.Printf("Elastic Query: %s\n", string(str))

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	res, err := m.client.Search(
		m.client.Search.WithContext(context.Background()),
		m.client.Search.WithIndex(changeIndex),
		m.client.Search.WithBody(&buf),
		m.client.Search.WithTrackTotalHits(true),
		m.client.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}

	totalHits := int(results["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits;",
		res.Status(),
		totalHits,
	)

	hits := results["hits"].(map[string]interface{})["hits"].([]interface{})

	var changes = make([]domain.SearchChange, totalHits)

	// Print the ID and document source for each hit.
	for i, hit := range hits {

		var changeDoc SearchChangeModel

		source := hit.(map[string]interface{})["_source"].(map[string]interface{})

		//log.Printf("source %s\n\n", source)

		err := encoders.Decode(source, &changeDoc)

		// just return if one fails
		if err != nil {
			return changes, err
		}

		changes[i] = m.toDomain(changeDoc)
	}

	return changes, err
}

func (m *ChangeSearchRepository) Create(change domain.SearchChange) (domain.SearchChange, error) {

	doc := m.toPersistence(change)

	bytes, err := json.Marshal(doc)

	req := esapi.IndexRequest{
		Index:      changeIndex,
		DocumentID: change.Id,
		Body:       strings.NewReader(string(bytes)),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), m.client)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fmt.Println(res, err)

	return change, err
}
