package use_cases

import (
	domain "diffme.dev/diffme-api/server/modules/organizations"
)

func (u *OrganizationUseCases) Update(orgId string, org domain.Organization) (domain.Organization, error) {
	return u.orgRepo.Update(orgId, org)
}
