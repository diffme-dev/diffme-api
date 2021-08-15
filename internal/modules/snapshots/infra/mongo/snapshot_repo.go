package mongo

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "snapshots"
)

type SnapshotModel struct {
	bongo.DocumentBase `bson:",inline"`
	ID                 string      `json:"id"`
	ReferenceID        string      `json:"reference_id"`
	Data               interface{} `json:"data"`
	Metadata           interface{} `json:"metadata"`
	Editor             string      `json:"id"`
	UpdatedAt          time.Time   `json:"updated_at"`
	CreatedAt          time.Time   `json:"created_at"`
}

type MongoSnapshotRepo struct {
	DB *bongo.Connection
}

func NewMongoSnapshotRepo(DB *bongo.Connection) domain.SnapshotRepo {
	return &MongoSnapshotRepo{DB}
}

func (m *MongoSnapshotRepo) _transformToDomain(doc SnapshotModel) domain.Snapshot {
	return domain.Snapshot{
		ID:          doc.ID,
		ReferenceID: doc.ReferenceID,
		Data:        doc.Data,
		Editor:      doc.Editor,
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *MongoSnapshotRepo) FindByID(id string) (snapshot domain.Snapshot, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err

}

func (m *MongoSnapshotRepo) FindByReferenceID(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindOne(bson.M{"reference_id": referenceID}, snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err
}

func (m *MongoSnapshotRepo) FindMostRecentByReference(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	query := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})
	err = query.Query.Sort("created_at").Limit(1).One(snapshotDoc)

	// TODO: how do you handle the null cases here?
	return m._transformToDomain(*snapshotDoc), err
}

func (m *MongoSnapshotRepo) FindForReference(referenceID string) (res []domain.Snapshot, err error) {

	result := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})

	page, err := result.Paginate(10, 10)

	snapshots := make([]domain.Snapshot, page.RecordsOnPage)

	for i := 0; i < page.RecordsOnPage; i++ {
		doc := &SnapshotModel{}
		_ = result.Next(doc)
		snapshots[i] = m._transformToDomain(*doc)
	}

	return snapshots, err
}

func (m *MongoSnapshotRepo) Create(params domain.CreateSnapshotParams) (res domain.Snapshot, err error) {

	snapshotDoc := &SnapshotModel{
		ID:          bson.NewObjectId().Hex(),
		ReferenceID: params.ID,
		Data:        params.Data,
		Metadata:    params.Metadata,
		Editor:      params.Editor,
		UpdatedAt:   time.Time{},
		CreatedAt:   params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err
}
