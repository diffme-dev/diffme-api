package mongo

import (
	"context"
	Infra "diffme.dev/diffme-api/internal/core/infra"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/wI2L/jsondiff"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var (
	modelName = "changes"
)

type ChangeRepo struct {
	DB *bongo.Connection
}

type Diff jsondiff.Operation

type ChangeModel struct {
	bongo.DocumentBase `bson:",inline"`
	ChangeSetId        string                 `bson:"change_set_id" json:"change_set_id"`
	ReferenceId        string                 `bson:"reference_id" json:"reference_id"`
	SnapshotId         string                 `bson:"snapshot_id" json:"snapshot_id"`
	Editor             string                 `bson:"editor" json:"editor"`
	Metadata           map[string]interface{} `bson:"metadata" json:"metadata"`
	Diff               Diff                   `bson:"diff" json:"diff"`
	UpdatedAt          time.Time              `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time              `bson:"created_at" json:"created_at"`
}

func NewMongoChangeRepo(DB *bongo.Connection) domain.ChangeRepository {
	return &ChangeRepo{DB: DB}
}

func (m *ChangeRepo) toDomain(doc ChangeModel) domain.Change {
	return domain.Change{
		Id:          doc.Id.Hex(),
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

func (m *ChangeRepo) toPersistence(change domain.Change) ChangeModel {
	return ChangeModel{
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

func (m *ChangeRepo) FindByID(id string) (snapshot domain.Change, err error) {
	objectID := bson.ObjectIdHex(id)
	changeDoc := &ChangeModel{}

	err = m.DB.Collection(modelName).FindById(objectID, changeDoc)

	return m.toDomain(*changeDoc), err

}

func (m *ChangeRepo) Create(change domain.Change) (res domain.Change, err error) {
	changeDoc := m.toPersistence(change)

	err = m.DB.Collection(modelName).Save(&changeDoc)

	if err != nil {
		fmt.Printf("error making change %s", err)
	}

	return m.toDomain(changeDoc), err
}

func (m *ChangeRepo) CreateMultiple(changes []domain.Change) (res []domain.Change, err error) {

	changeDocs := make([]ChangeModel, len(changes))

	for i, change := range changes {

		changeDoc := m.toPersistence(change)

		err := m.DB.Collection(modelName).Save(&changeDoc)

		if err != nil {
			println(err)
			continue
		}

		changeDocs[i] = changeDoc
	}

	// transform back
	newChanges := make([]domain.Change, len(changeDocs))

	for i, changeDoc := range changeDocs {
		newChanges[i] = m.toDomain(changeDoc)
	}

	return newChanges, err
}

func (m *ChangeRepo) CreateMultipleTxn(changes []domain.Change) (res []domain.Change, err error) {
	println("starting transaction")

	changeDocs := make([]ChangeModel, len(changes))

	for i, change := range changes {
		changeDocs[i] = m.toPersistence(change)
	}

	client, err := Infra.NewMongoConnection()
	changesCollection := client.Database("diffme").Collection(modelName)
	session, err := client.StartSession()

	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(); err != nil {
			return err
		}

		for _, doc := range changeDocs {
			_, err := changesCollection.InsertOne(
				sessionContext,
				doc,
			)

			if err != nil {
				return err
			}
		}

		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			log.Printf("error %s", abortErr)
			//panic(abortErr)
		}
		log.Printf("error %s", err)
		//panic(err)
	}

	// transform back
	newChanges := make([]domain.Change, len(changeDocs))

	for i, changeDoc := range changeDocs {
		newChanges[i] = m.toDomain(changeDoc)
	}

	return newChanges, err
}
