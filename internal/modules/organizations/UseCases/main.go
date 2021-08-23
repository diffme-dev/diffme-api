package UseCases

import (
	OrganizationDomain "diffme.dev/diffme-api/internal/modules/organizations"
)

type OrganizationUseCases struct {
	orgRepo OrganizationDomain.OrganizationRepository
}

func NewOrganizationUseCases(orgRepo OrganizationDomain.OrganizationRepository) OrganizationDomain.OrganizationUseCases {
	return &OrganizationUseCases{
		orgRepo: orgRepo,
	}
}