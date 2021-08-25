package mongo

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
	"fmt"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "users"
)

type UserAuthProviderModel struct {
	Provider       string `bson:"provider" json:"provider"`
	ProviderUserId string `bson:"provider_user_id" json:"provider_user_id""`
}

type UserModel struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string                `bson:"name" json:"name"`
	FirstName          string                `bson:"first_name" json:"first_name"`
	LastName           string                `bson:"last_name" json:"last_name"`
	PhoneNumber        string                `bson:"phone_number" json:"phone_number"`
	Email              string                `bson:"email" json:"email"`
	ProfileUrl         string                `bson:"profile_url" json:"profile_url"`
	Auth               UserAuthProviderModel `bson:"auth" json:"auth"`
	UpdatedAt          time.Time             `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time             `bson:"created_at" json:"created_at"`
}

type UserRepo struct {
	DB *bongo.Connection
}

func NewMongoUserRepo(DB *bongo.Connection) domain.UserRepository {
	return &UserRepo{DB: DB}
}

func (m *UserRepo) toDomain(doc *UserModel) *domain.User {
	return &domain.User{
		Id:          doc.Id.Hex(),
		Name:        doc.Name,
		FirstName:   doc.FirstName,
		LastName:    doc.LastName,
		PhoneNumber: doc.PhoneNumber,
		Email:       doc.Email,
		ProfileUrl:  doc.ProfileUrl,
		Auth:        domain.UserAuthProvider(doc.Auth),
		UpdatedAt:   doc.UpdatedAt,
		CreatedAt:   doc.CreatedAt,
	}
}

func (m *UserRepo) FindById(id string) (snapshot *domain.User, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &UserModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *UserRepo) Create(params domain.CreateUserParams) (res *domain.User, err error) {

	auth := UserAuthProviderModel(*params.Auth)

	fmt.Printf("\nAuth: %+v\n", auth)

	snapshotDoc := &UserModel{
		Name:        params.Name,
		LastName:    params.LastName,
		FirstName:   params.FirstName,
		ProfileUrl:  params.ProfileUrl,
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
		Auth:        auth,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	fmt.Printf("\n\nNew User: %+v\n\n", snapshotDoc)

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	if err != nil {
		return nil, err
	}

	return m.toDomain(snapshotDoc), err
}

// FIXME:
func (m *UserRepo) Update(userId string, params domain.User) (res *domain.User, err error) {

	snapshotDoc := &UserModel{
		UpdatedAt: time.Now(),
		CreatedAt: params.CreatedAt,
	}

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	return m.toDomain(snapshotDoc), err
}
