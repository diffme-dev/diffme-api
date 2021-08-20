package organizations

import (
	"time"
)

type Organization struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type OrganizationRepository interface {
	FindById(id string) (Organization, error)
	Update(orgId string, org Organization) (Organization, error)
	Create(org Organization) (Organization, error)
}

type OrganizationUseCases interface {
}
