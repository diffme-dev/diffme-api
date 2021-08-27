package team_members

import "time"

type TeamMember struct {
	Id             string    `json:"id"`
	UserId         string    `json:"user_id"`
	OrganizationId string    `json:"organization_id"`
	ACL            string    `json:"acl"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type TeamMemberRepository interface {
	FindById(id string) (*TeamMember, error)
	Create(member TeamMember) (*TeamMember, error)
}

type TeamMemberUseCases interface {
	InviteTeamMember(params TeamMember) (*TeamMember, error)
}
