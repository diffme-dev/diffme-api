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

func (m *MongoChangeRepo) toDomain(doc ChangeModel) domain.Change {
	return domain.Change{
		ID:        doc.ID,
		Editor:    doc.Editor,
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (m *MongoChangeRepo) toPersistence(change domain.Change) ChangeModel {
	return ChangeModel{
		ID:        change.ID,
		Editor:    change.Editor,
		UpdatedAt: change.UpdatedAt,
		CreatedAt: change.CreatedAt,
	}
}

func (m *MongoChangeRepo) FindByID(id string) (snapshot domain.Change, err error) {
	objectID := bson.ObjectIdHex(id)
	changeDoc := &ChangeModel{}

	err = m.DB.Collection(modelName).FindById(objectID, changeDoc)

	return m.toDomain(*changeDoc), err

}

func (m *MongoChangeRepo) Create(change domain.Change) (res domain.Change, err error) {

	changeDoc := m.toPersistence(change)

	err = m.DB.Collection(modelName).Save(&changeDoc)

	return m.toDomain(changeDoc), err
}

func (m *MongoChangeRepo) CreateMultiple(changes []domain.Change) (res []domain.Change, err error) {

	changeDocs := make([]ChangeModel, len(changes))

	for _, change := range changes {
		changeDocs = append(changeDocs, m.toPersistence(change))
	}

	session := m.DB.Session

	err = session.DB(m.DB.Config.Database).C(modelName).Insert(changeDocs)

	if err != nil {
		return nil, err
	}

	// transform back
	newChanges := make([]domain.Change, len(changeDocs))

	for _, changeDoc := range changeDocs {
		newChanges = append(newChanges, m.toDomain(changeDoc))
	}

	return newChanges, err
}
