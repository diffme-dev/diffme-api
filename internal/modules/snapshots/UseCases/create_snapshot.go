package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"diffme.dev/diffme-api/internal/modules/snapshots/services"
	"fmt"
)

func (u *SnapshotUseCases) CreateSnapshot(snapshotParams domain.CreateSnapshotParams) (domain.Snapshot, error) {
	// Note: the ID on this is actually what we consider a reference id
	referenceId := snapshotParams.Id
	lastSnapshot, err := u.snapshotRepo.FindMostRecentByReference(referenceId)
	snapshot, err := u.snapshotRepo.Create(snapshotParams)

	fmt.Printf("Latest Ref: %s\n", lastSnapshot)

	services.SnapshotCreated(lastSnapshot, snapshot)

	return snapshot, err
}
