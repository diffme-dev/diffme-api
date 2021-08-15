package elasticsearch

import (
	"bytes"
	"context"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"log"
	"strings"
	"time"
)

type ElasticChangeSearchRepository struct {
	client *elasticsearch.Client
}

type SearchChangeModel struct {
	ID          string    `json:"id"`
	ChangeSetID string    `json:"change_set_id"`
	ReferenceID string    `json:"reference_id"`
	Editor      string    `json:"id"`
	Metadata    []byte    `json:"metadata"`
	Diffs       []byte    `json:"diffs"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewElasticSearchChangeRepo(client *elasticsearch.Client) domain.SearchChangeRepository {
	return &ElasticChangeSearchRepository{client: client}
}

func (m *ElasticChangeSearchRepository) Query(id string) (snapshot domain.Change, err error) {

	var (
		results map[string]interface{}
		buf     bytes.Buffer
	)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := m.client.Search(
		m.client.Search.WithContext(context.Background()),
		m.client.Search.WithIndex("test"),
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
	// Print the ID and document source for each hit.
	for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))

	return domain.Change{}, err
}

func (m *ElasticChangeSearchRepository) Create(change domain.Change) (res domain.Change, err error) {

	//changeDoc := &ChangeModel{
	//	ID:                 change.ID,
	//	ReferenceID:        change.ReferenceID,
	//	ChangeSetID:        change.ChangeSetID,
	//	Editor:             change.Editor,
	//	Metadata:           change.Metadata,
	//	Diffs:              change.Diffs,
	//	PreviousSnapshotID: change.PreviousSnapshotID,
	//}
	//
	//err = m.DB.Collection(modelName).Save(changeDoc)
	return domain.Change{}, err
}
