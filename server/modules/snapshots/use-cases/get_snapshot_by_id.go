package use_cases

import (
	"diffme.dev/diffme-api/server/modules/snapshots"
)

func (u *SnapshotUseCases) GetSnapshotByID(id string) (*domain.Snapshot, error) {
	return u.snapshotRepo.FindByID(id)
}
