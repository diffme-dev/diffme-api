package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) SearchChange(match domain.SearchRequest) ([]domain.SearchChange, error) {

	searchChanges, err := a.searchChangeRepo.Query(match)

	if err != nil {
		return make([]domain.SearchChange, 0), err
	}

	return searchChanges, nil
}
