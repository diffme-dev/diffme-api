package mongo

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "users"
)

type UserModel struct {
	bongo.DocumentBase `bson:",inline"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
}

type UserRepo struct {
	DB *bongo.Connection
}

func NewMongoUserRepo(DB *bongo.Connection) domain.UserRepository {
	return &UserRepo{DB: DB}
}

func (m *UserRepo) toDomain(doc *UserModel) domain.User {
	if doc == nil {
		return domain.User{}
	}

	return domain.User{
		Id:        doc.Id.Hex(),
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (m *UserRepo) FindById(id string) (snapshot domain.User, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &UserModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *UserRepo) Create(params domain.User) (res domain.User, err error) {

	snapshotDoc := &UserModel{
		UpdatedAt: time.Now(),
		CreatedAt: params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}

// TODO: fix this...
func (m *UserRepo) Update(userId string, params domain.User) (res domain.User, err error) {

	snapshotDoc := &UserModel{
		UpdatedAt: time.Now(),
		CreatedAt: params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}
