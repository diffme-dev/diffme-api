package mongo

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"fmt"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "changes"
)

type ChangeRepo struct {
	DB *bongo.Connection
}

type ChangeDiffModel struct {
	Type     string                   `json:"op" bson:"type"`
	From     interfaces.StringPointer `json:"from,omitempty" bson:"from"`
	Path     interfaces.StringPointer `json:"path" bson:"path"`
	Value    interface{}              `json:"value,omitempty" bson:"value"`
	OldValue interface{}              `json:"old_value,omitempty" bson:"old_value"`
}

type ChangeModel struct {
	bongo.DocumentBase `bson:",inline"`
	Label              *string                `bson:"label" json:"label"`
	EventName          *string                `bson:"event_name" json:"event_name"`
	ChangeSetId        string                 `bson:"change_set_id" json:"change_set_id"`
	ReferenceId        string                 `bson:"reference_id" json:"reference_id"`
	SnapshotId         string                 `bson:"snapshot_id" json:"snapshot_id"`
	Editor             string                 `bson:"editor" json:"editor"`
	Metadata           map[string]interface{} `bson:"metadata" json:"metadata"`
	Diff               ChangeDiffModel        `bson:"diff" json:"diff"`
	UpdatedAt          time.Time              `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time              `bson:"created_at" json:"created_at"`
}

func NewMongoChangeRepo(DB *bongo.Connection) domain.ChangeRepository {
	return &ChangeRepo{DB: DB}
}

func (m *ChangeRepo) toDomain(doc ChangeModel) domain.Change {
	return domain.Change{
		Id:          doc.Id.Hex(),
		Label:       doc.Label,
		EventName:   doc.EventName,
		ChangeSetId: doc.ChangeSetId,
		ReferenceId: doc.ReferenceId,
		SnapshotId:  doc.SnapshotId,
		Editor:      doc.Editor,
		Metadata:    doc.Metadata,
		Diff: domain.ChangeDiff{
			Type:     doc.Diff.Type,
			OldValue: doc.Diff.OldValue,
			Value:    doc.Diff.Value,
			From:     doc.Diff.From,
			Path:     doc.Diff.Path,
		},
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (m *ChangeRepo) toPersistence(change domain.Change) ChangeModel {
	return ChangeModel{
		Label:       change.Label,
		EventName:   change.EventName,
		ChangeSetId: change.ChangeSetId,
		ReferenceId: change.ReferenceId,
		SnapshotId:  change.SnapshotId,
		Editor:      change.Editor,
		Metadata:    change.Metadata,
		Diff: ChangeDiffModel{
			OldValue: change.Diff.OldValue,
			Type:     change.Diff.Type,
			Value:    change.Diff.Value,
			From:     change.Diff.From,
			Path:     change.Diff.Path,
		},
		UpdatedAt: change.UpdatedAt,
		CreatedAt: change.CreatedAt,
	}
}

func (m *ChangeRepo) FindById(id string) (snapshot domain.Change, err error) {
	objectID := bson.ObjectIdHex(id)
	changeDoc := &ChangeModel{}

	err = m.DB.Collection(modelName).FindById(objectID, changeDoc)

	return m.toDomain(*changeDoc), err

}

func (m *ChangeRepo) FindByReferenceId(referenceId string) (snapshot []domain.Change, err error) {

	result := m.DB.Collection(modelName).Find(bson.M{"reference_id": referenceId})
	page, err := result.Paginate(10, 0)

	if err != nil {
		return []domain.Change{}, err
	}

	changes := make([]domain.Change, page.RecordsOnPage)

	for i := 0; i < page.RecordsOnPage; i++ {
		doc := &ChangeModel{}
		_ = result.Next(doc)
		changes[i] = m.toDomain(*doc)
	}

	return changes, err

}

func (m *ChangeRepo) Find(query domain.QueryChangesRequest) (snapshot []domain.Change, err error) {

	findQuery := bson.M{}

	// set before/after for the query
	if query.Before != nil || query.After != nil {
		idQuery := bson.M{}
		if query.Before != nil {
			idQuery["lt"] = query.Before
		}
		if query.After != nil {
			idQuery["gte"] = query.After
		}
		findQuery = bson.M{
			"_id": idQuery,
		}
	}

	// update the limit if it is nil
	if query.Limit == nil {
		l := 50
		query.Limit = &l
	}

	if query.Sort == nil {
		f := "-created_at"
		query.Sort = &f
	}

	fmt.Printf("Limit: %d. Sort: %s\n", *query.Limit, *query.Sort)

	var changeDocs []ChangeModel

	err = m.DB.Collection(modelName).Find(findQuery).Query.Limit(*query.Limit).Sort(*query.Sort).All(&changeDocs)

	if err != nil {
		fmt.Printf("Error %s\n", err)
		return []domain.Change{}, err
	}

	fmt.Printf("found %d changes\n", len(changeDocs))

	changes := make([]domain.Change, len(changeDocs))

	for i, doc := range changeDocs {
		changes[i] = m.toDomain(doc)
	}

	return changes, err

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
