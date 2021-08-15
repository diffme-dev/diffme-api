package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
)

func (u *SnapshotUseCases) FindMostRecentReference(id string) (domain.Snapshot, error) {
	return u.snapshotRepo.FindMostRecentByReference(id)
}
