package use_cases

import (
	domain "diffme.dev/diffme-api/server/modules/organizations"
)

func (u *OrganizationUseCases) Create(org domain.Organization) (domain.Organization, error) {
	return u.orgRepo.Create(org)
}
