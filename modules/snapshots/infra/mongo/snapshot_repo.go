package mongo

import (
	domain "diffme.dev/diffme-api/modules/snapshots"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "snapshots"
)

type SnapshotModel struct {
	bongo.DocumentBase `bson:",inline"`
	ID                 string    `json:"id" validate:"required"`
	ReferenceID        string    `json:"reference_id" validate:"required"`
	Data               []byte    `json:"data" validate:"required"`
	Editor             string    `json:"id" validate:"required"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}

type mongoSnapshotRepo struct {
	DB *bongo.Connection
}

func NewMongoSnapshotRepo(DB *bongo.Connection) domain.SnapshotRepo {
	return &mongoSnapshotRepo{DB}
}

func (m *mongoSnapshotRepo) _transformToDomain(doc SnapshotModel) domain.Snapshot {
	return domain.Snapshot{
		ID:          doc.ID,
		ReferenceID: doc.ReferenceID,
		Data:        doc.Data,
		Editor:      doc.Editor,
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *mongoSnapshotRepo) FindByID(id string) (snapshot domain.Snapshot, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err

}

func (m *mongoSnapshotRepo) FindByReferenceID(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindOne(bson.M{"reference_id": referenceID}, snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err
}

func (m *mongoSnapshotRepo) FindMostRecentByReference(referenceID string) (res domain.Snapshot, err error) {
	snapshotDoc := &SnapshotModel{}

	query := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})
	err = query.Query.Sort("created_at").Limit(1).One(snapshotDoc)

	// TODO: how do you handle the null cases here?
	return m._transformToDomain(*snapshotDoc), err
}

func (m *mongoSnapshotRepo) FindSnapshotsForReference(referenceID string) (res []domain.Snapshot, err error) {

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

func (m *mongoSnapshotRepo) CreateSnapshot(params domain.CreateSnapshotParams) (res domain.Snapshot, err error) {

	snapshotDoc := &SnapshotModel{
		ID:          params.ID,
		ReferenceID: params.ReferenceID,
		Data:        params.Data,
		Editor:      params.Editor,
		UpdatedAt:   params.UpdatedAt,
		CreatedAt:   params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m._transformToDomain(*snapshotDoc), err
}
