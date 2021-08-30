package use_cases

import (
	"diffme.dev/diffme-api/server/modules/changes"
)

func (a *ChangeUseCases) GetChanges(query domain.QueryChangesRequest) ([]domain.Change, error) {

	changes, err := a.changeRepo.Find(query)

	if err != nil {
		return make([]domain.Change, 0), err
	}

	return changes, nil
}
