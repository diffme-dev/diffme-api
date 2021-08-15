package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) SearchChange(query string) ([]domain.Change, error) {

	changes := make([]domain.Change, 5)

	return changes, nil
}
