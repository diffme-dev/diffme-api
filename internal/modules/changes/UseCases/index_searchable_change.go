package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) IndexSearchableChange(change domain.Change) (domain.Change, error) {
	return domain.Change{}, nil
}
