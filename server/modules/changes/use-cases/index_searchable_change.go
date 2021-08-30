package use_cases

import (
	"diffme.dev/diffme-api/server/modules/changes"
)

func (a *ChangeUseCases) IndexSearchableChange(change domain.Change) (domain.SearchChange, error) {
	searchChange := domain.SearchChange{
		ChangeSetId: change.ChangeSetId,
		Id:          change.Id,
		Editor:      change.Editor,
		Diff:        change.Diff,
		ReferenceId: change.ReferenceId,
		SnapshotId:  change.SnapshotId,
		UpdatedAt:   change.UpdatedAt,
		CreatedAt:   change.CreatedAt,
	}

	_, err := a.searchChangeRepo.Create(searchChange)

	return searchChange, err
}
