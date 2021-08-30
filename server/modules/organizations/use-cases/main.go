package use_cases

import (
	OrganizationDomain "diffme.dev/diffme-api/server/modules/organizations"
)

type OrganizationUseCases struct {
	orgRepo OrganizationDomain.OrganizationRepository
}

func NewOrganizationUseCases(orgRepo OrganizationDomain.OrganizationRepository) OrganizationDomain.OrganizationUseCases {
	return &OrganizationUseCases{
		orgRepo: orgRepo,
	}
}
