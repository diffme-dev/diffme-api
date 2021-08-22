package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"time"
)

func (u *SnapshotUseCases) FindMostRecentReference(id string, before *time.Time) (domain.Snapshot, error) {
	return u.snapshotRepo.FindMostRecentByReference(id, before)
}
