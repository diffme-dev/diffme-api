package UseCases

import OrganizationDomain "diffme.dev/diffme-api/internal/modules/organizations"

func (u *OrganizationUseCases) GetOrganization(orgId string) (OrganizationDomain.Organization, error) {
	return u.orgRepo.FindById(orgId)
}
