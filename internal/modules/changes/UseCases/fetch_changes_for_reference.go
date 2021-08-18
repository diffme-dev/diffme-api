package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) FetchChangeForReferenceId(referenceID string) ([]domain.Change, error) {
	return a.changeRepo.FindByReferenceId(referenceID)
}
