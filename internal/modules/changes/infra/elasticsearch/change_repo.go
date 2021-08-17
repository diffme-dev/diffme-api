package elasticsearch

import (
	"bytes"
	"context"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/mitchellh/mapstructure"
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
	Id          string                 `json:"id"`
	ChangeSetId string                 `json:"change_set_id"`
	SnapshotId  string                 `json:"snapshot_id"`
	ReferenceId string                 `json:"reference_id"`
	Editor      string                 `json:"id"`
	Metadata    map[string]interface{} `json:"metadata"`
	Diff        Diff                   `json:"diff"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CreatedAt   time.Time              `json:"created_at"`
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

func (m *ChangeSearchRepository) Query(match map[string]interface{}) ([]domain.SearchChange, error) {

	var (
		results map[string]interface{}
		buf     bytes.Buffer
	)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": match,
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := m.client.Search(
		m.client.Search.WithContext(context.Background()),
		m.client.Search.WithIndex(changeIndex),
		m.client.Search.WithBody(&buf),
		m.client.Search.WithTrackTotalHits(true),
		m.client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(results["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(results["took"].(float64)),
	)

	var changes = make([]domain.SearchChange, len(results))

	// Print the ID and document source for each hit.
	for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {

		var changeDoc SearchChangeModel
		err := mapstructure.Decode(hit, &changeDoc)

		if err != nil {
			println(err)
			continue
		}

		changes = append(changes, m.toDomain(changeDoc))
	}

	return changes, err
}

func (m *ChangeSearchRepository) Create(change domain.SearchChange) (domain.SearchChange, error) {

	bytes, err := json.Marshal(m.toPersistence(change))

	res, err := m.client.Index(
		changeIndex,
		strings.NewReader(string(bytes)),
		m.client.Index.WithDocumentID(change.Id))

	fmt.Println("Elastic Search:")
	fmt.Println(res, err)

	return change, err
}
