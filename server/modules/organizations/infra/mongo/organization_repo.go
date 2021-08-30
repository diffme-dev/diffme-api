package mongo

import (
	domain "diffme.dev/diffme-api/server/modules/organizations"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "organizations"
)

type OrganizationModel struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string    `bson:"name" json:"name"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
}

type OrganizationRepo struct {
	DB *bongo.Connection
}

func NewMongoOrganizationRepo(DB *bongo.Connection) domain.OrganizationRepository {
	return &OrganizationRepo{DB: DB}
}

func (m *OrganizationRepo) toDomain(doc *OrganizationModel) domain.Organization {
	if doc == nil {
		return domain.Organization{}
	}

	return domain.Organization{
		Id:        doc.Id.Hex(),
		Name:      doc.Name,
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (m *OrganizationRepo) FindById(id string) (snapshot domain.Organization, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &OrganizationModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *OrganizationRepo) Create(params domain.Organization) (res domain.Organization, err error) {

	snapshotDoc := &OrganizationModel{
		Name:      params.Name,
		UpdatedAt: time.Now(),
		CreatedAt: params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}

func (m *OrganizationRepo) Update(userId string, params domain.Organization) (res domain.Organization, err error) {
	snapshotDoc := &OrganizationModel{
		Name:      params.Name,
		UpdatedAt: time.Now(),
		CreatedAt: params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}
