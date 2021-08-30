package use_cases

import (
	OrganizationDomain "diffme.dev/diffme-api/server/modules/organizations"
)

func (u *OrganizationUseCases) GetOrganization(orgId string) (OrganizationDomain.Organization, error) {
	return u.orgRepo.FindById(orgId)
}
