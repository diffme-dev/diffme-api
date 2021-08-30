package mongo

import (
	"diffme.dev/diffme-api/server/modules/snapshots"
	"diffme.dev/diffme-api/server/shared"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "snapshots"
)

type SnapshotModel struct {
	bongo.DocumentBase `bson:",inline"`
	Label              *string                `bson:"label" json:"label"`
	EventName          *string                `bson:"event_name" json:"event_name"`
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

func (m *SnapshotRepo) toDomain(doc *SnapshotModel) *domain.Snapshot {
	if doc == nil {
		return nil
	}

	return &domain.Snapshot{
		Id:          doc.Id.Hex(),
		Label:       doc.Label,
		EventName:   doc.EventName,
		ReferenceId: doc.ReferenceId,
		Data:        doc.Data,
		Editor:      doc.Editor,
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *SnapshotRepo) toPersistence(doc domain.Snapshot) *SnapshotModel {
	return &SnapshotModel{
		Label:       doc.Label,
		EventName:   doc.EventName,
		ReferenceId: doc.ReferenceId,
		Data:        doc.Data,
		Editor:      doc.Editor,
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *SnapshotRepo) FindByID(id string) (snapshot *domain.Snapshot, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &SnapshotModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *SnapshotRepo) FindByReferenceID(referenceID string) (*domain.Snapshot, error) {
	snapshotDoc := &SnapshotModel{}

	err := m.DB.Collection(modelName).FindOne(bson.M{"reference_id": referenceID}, snapshotDoc)

	return m.toDomain(snapshotDoc), err
}

func (m *SnapshotRepo) FindMostRecentByReference(referenceId string, before *time.Time) (*domain.Snapshot, error) {
	snapshotDoc := &SnapshotModel{}

	dbQuery := bson.M{"reference_id": referenceId}

	if before != nil {
		dbQuery["created_at"] = bson.M{
			"$lte": before,
		}
	}

	//shared.Logger.Infof("mongo query %+v", dbQuery)

	err := m.DB.Collection(modelName).Find(dbQuery).Query.Sort("-created_at").Limit(1).One(&snapshotDoc)

	if err != nil {
		shared.GetSugarLogger().Errorf("error occured %v", err)

		return nil, err
	}

	//shared.Logger.Infof("mongo result %+v", snapshotDoc)

	return m.toDomain(snapshotDoc), nil
}

func (m *SnapshotRepo) FindForReference(referenceID string) (res []domain.Snapshot, err error) {

	result := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceID})

	page, err := result.Paginate(10, 0)

	snapshots := make([]domain.Snapshot, page.RecordsOnPage)

	for i := 0; i < page.RecordsOnPage; i++ {
		doc := &SnapshotModel{}
		_ = result.Next(doc)
		snapshots[i] = *m.toDomain(doc)
	}

	return snapshots, err
}

func (m *SnapshotRepo) Create(params domain.CreateSnapshotParams) (*domain.Snapshot, error) {

	snapshotDoc := m.toPersistence(domain.Snapshot{
		ReferenceId: params.Id,
		Label:       params.Label,
		EventName:   params.EventName,
		Data:        params.Data,
		Metadata:    params.Metadata,
		Editor:      params.Editor,
		UpdatedAt:   time.Now(),
		CreatedAt:   params.CreatedAt,
	})

	err := m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}
