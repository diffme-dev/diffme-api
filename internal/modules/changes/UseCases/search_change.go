package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) SearchChange(query string) ([]domain.SearchChange, error) {

	changes := make([]domain.SearchChange, 5)

	return changes, nil
}
