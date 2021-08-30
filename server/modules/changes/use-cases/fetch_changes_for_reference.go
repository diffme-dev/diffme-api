package use_cases

import (
	"diffme.dev/diffme-api/server/modules/changes"
)

func (a *ChangeUseCases) FetchChangeForReferenceId(referenceID string) ([]domain.Change, error) {
	return a.changeRepo.FindByReferenceId(referenceID)
}
