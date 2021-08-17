package mongo

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"fmt"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "snapshots"
)

type SnapshotModel struct {
	bongo.DocumentBase `bson:",inline"`
	ReferenceId        string                 `bson:"reference_id" json:"reference_id"`
	Data               map[string]interface{} `bson:"data" json:"data"`
	Metadata           map[string]interface{} `bson:"metadata" json:"metadata"`
	Editor             string                 `bson:"editor" json:"id"`
	UpdatedAt          time.Time              `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time              `bson:"created_at" json:"created_at"`
}

type SnapshotRepo struct {
	DB *bongo.Connection
}

func NewMongoSnapshotRepo(DB *bongo.Connection) domain.SnapshotRepo {
	return &SnapshotRepo{DB: DB}
}

func (m *SnapshotRepo) toDomain(doc SnapshotModel) domain.Snapshot {
	return domain.Snapshot{
		Id:          doc.Id.Hex(),
		ReferenceId: doc.ReferenceId,
		Data:        doc.Data,
		Editor:      doc.Editor,
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *SnapshotRepo) FindByID(id string) (snapshot domain.Snapshot, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(*snapshotDoc), err

}

func (m *SnapshotRepo) FindByReferenceID(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindOne(bson.M{"reference_id": referenceID}, snapshotDoc)

	return m.toDomain(*snapshotDoc), err
}

func (m *SnapshotRepo) FindMostRecentByReference(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	query := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})
	err = query.Query.Sort("created_at").Limit(1).One(snapshotDoc)

	// TODO: how do you handle the null cases here?
	return m.toDomain(*snapshotDoc), err
}

func (m *SnapshotRepo) FindForReference(referenceID string) (res []domain.Snapshot, err error) {

	result := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})

	page, err := result.Paginate(10, 0)

	snapshots := make([]domain.Snapshot, page.RecordsOnPage)

	for i := 0; i < page.RecordsOnPage; i++ {
		doc := &SnapshotModel{}
		_ = result.Next(doc)
		snapshots[i] = m.toDomain(*doc)
	}

	return snapshots, err
}

func (m *SnapshotRepo) Create(params domain.CreateSnapshotParams) (res domain.Snapshot, err error) {

	snapshotDoc := &SnapshotModel{
		ReferenceId: params.Id,
		Data:        params.Data,
		Metadata:    params.Metadata,
		Editor:      params.Editor,
		UpdatedAt:   time.Now(),
		CreatedAt:   params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	fmt.Printf("SNAP %s", m.toDomain(*snapshotDoc))

	return m.toDomain(*snapshotDoc), err
}
