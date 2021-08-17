package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
)

func (a *ChangeUseCases) IndexSearchableChange(change domain.Change) (domain.SearchChange, error) {
	searchChange := domain.SearchChange{
		ChangeSetId: change.ChangeSetId,
		Id:          change.Id,
		Diff:        change.Diff,
		ReferenceId: change.ReferenceId,
		SnapshotId:  change.SnapshotId,
		UpdatedAt:   change.UpdatedAt,
		CreatedAt:   change.CreatedAt,
	}

	_, err := a.searchChangeRepo.Create(searchChange)

	return searchChange, err
}
