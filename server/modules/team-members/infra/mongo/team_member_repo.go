package mongo

import (
	domain "diffme.dev/diffme-api/server/modules/team-members"
	"fmt"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	modelName = "team_members"
)

type TeamMemberModel struct {
	bongo.DocumentBase `bson:",inline"`
	UserId             string    `bson:"user_id" json:"user_id"`
	OrganizationId     string    `bson:"organization_id" json:"organization_id"`
	ACL                string    `bson:"acl" json:"acl"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
}

type TeamMemberRepo struct {
	DB *bongo.Connection
}

func NewMongoTeamMemberRepo(DB *bongo.Connection) domain.TeamMemberRepository {
	return &TeamMemberRepo{DB: DB}
}

func (m *TeamMemberRepo) toDomain(doc *TeamMemberModel) *domain.TeamMember {
	return &domain.TeamMember{
		Id:             doc.Id.Hex(),
		UserId:         doc.UserId,
		OrganizationId: doc.OrganizationId,
		ACL:            doc.ACL,
		UpdatedAt:      doc.UpdatedAt,
		CreatedAt:      doc.CreatedAt,
	}
}

func (m *TeamMemberRepo) toPersistence(doc *domain.TeamMember) *TeamMemberModel {
	return &TeamMemberModel{
		UserId:         doc.UserId,
		OrganizationId: doc.OrganizationId,
		ACL:            doc.ACL,
		UpdatedAt:      doc.UpdatedAt,
		CreatedAt:      doc.CreatedAt,
	}
}

func (m *TeamMemberRepo) FindById(id string) (snapshot *domain.TeamMember, err error) {
	objectID := bson.ObjectIdHex(id)
	snapshotDoc := &TeamMemberModel{}

	err = m.DB.Collection(modelName).FindById(objectID, snapshotDoc)

	return m.toDomain(snapshotDoc), err

}

func (m *TeamMemberRepo) Create(params domain.TeamMember) (res *domain.TeamMember, err error) {

	snapshotDoc := m.toPersistence(&params)

	fmt.Printf("\n\nNew User: %+v\n\n", snapshotDoc)

	err = m.DB.Collection(modelName).Save(snapshotDoc)

	if err != nil {
		return nil, err
	}

	return m.toDomain(snapshotDoc), err
}
