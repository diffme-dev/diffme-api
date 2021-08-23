package UseCases

import domain "diffme.dev/diffme-api/internal/modules/organizations"

func (u *OrganizationUseCases) Create(org domain.Organization) (domain.Organization, error) {
	return u.orgRepo.Create(org)
}
