package mongo

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "changes"
)

type MongoChangeRepo struct {
	DB *bongo.Connection
}

type ChangeModel struct {
	bongo.DocumentBase `bson:",inline"`
	ID                 string    `json:"id"`
	ChangeSetID        string    `json:"change_set_id"`
	ReferenceID        string    `json:"reference_id"`
	PreviousSnapshotID string    `json:"previous_snapshot_id"`
	CurrentSnapshotID  string    `json:"current_snapshot_id"`
	Editor             string    `json:"id"`
	Metadata           []byte    `json:"metadata"`
	Diffs              []byte    `json:"diffs"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewMongoChangeRepo(DB *bongo.Connection) domain.ChangeRepository {
	return &MongoChangeRepo{DB: DB}
}

func (m *MongoChangeRepo) _transformToDomain(doc ChangeModel) domain.Change {
	return domain.Change{
		ID:        doc.ID,
		Editor:    doc.Editor,
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (m *MongoChangeRepo) FindByID(id string) (snapshot domain.Change, err error) {
	objectID := bson.ObjectIdHex(id)
	changeDoc := &ChangeModel{}

	err = m.DB.Collection(modelName).FindById(objectID, changeDoc)

	return m._transformToDomain(*changeDoc), err

}

func (m *MongoChangeRepo) Create(change domain.Change) (res domain.Change, err error) {

	changeDoc := &ChangeModel{
		ID:                 change.ID,
		ReferenceID:        change.ReferenceID,
		ChangeSetID:        change.ChangeSetID,
		Editor:             change.Editor,
		Metadata:           change.Metadata,
		Diffs:              change.Diffs,
		PreviousSnapshotID: change.PreviousSnapshotID,
	}

	err = m.DB.Collection(modelName).Save(changeDoc)

	return m._transformToDomain(*changeDoc), err
}
