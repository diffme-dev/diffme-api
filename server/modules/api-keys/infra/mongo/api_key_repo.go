package mongo

import (
	domain "diffme.dev/diffme-api/server/modules/api-keys"
	"fmt"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "api_keys"
)

type ApiKeyModel struct {
	bongo.DocumentBase `bson:",inline"`
	OrganizationId     string    `bson:"organization_id" json:"organization_id"`
	Label              string    `bson:"label" json:"label"`
	ApiKey             string    `bson:"api_key" json:"api_key"`
	ACL                string    `bson:"acl" json:"acl"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
}

type ApiKeyRepo struct {
	DB *bongo.Connection
}

func NewMongoApiKeyRepo(DB *bongo.Connection) domain.ApiKeyRepository {
	return &ApiKeyRepo{DB: DB}
}

func (m *ApiKeyRepo) toDomain(doc *ApiKeyModel) *domain.ApiKey {
	return &domain.ApiKey{
		Id:             doc.Id.Hex(),
		Label:          doc.Label,
		ApiKey:         doc.ApiKey,
		OrganizationId: doc.OrganizationId,
		UpdatedAt:      doc.UpdatedAt,
		CreatedAt:      doc.CreatedAt,
	}
}

func (m *ApiKeyRepo) toPersistence(doc *domain.ApiKey) *ApiKeyModel {
	return &ApiKeyModel{
		Label:          doc.Label,
		OrganizationId: doc.OrganizationId,
		ApiKey:         doc.ApiKey,
		ACL:            doc.ACL,
		UpdatedAt:      doc.UpdatedAt,
		CreatedAt:      doc.CreatedAt,
	}
}

func (m *ApiKeyRepo) FindById(id string) (snapshot *domain.ApiKey, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &ApiKeyModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *ApiKeyRepo) Create(params domain.ApiKey) (res *domain.ApiKey, err error) {

	snapshotDoc := m.toPersistence(&params)

	fmt.Printf("\n\nNew Api Key: %+v\n\n", snapshotDoc)

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	if err != nil {
		return nil, err
	}

	return m.toDomain(snapshotDoc), err
}
