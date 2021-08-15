package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
)

func (u *SnapshotUseCases) GetSnapshotByID(id string) (domain.Snapshot, error) {
	return u.snapshotRepo.FindByID(id)
}
