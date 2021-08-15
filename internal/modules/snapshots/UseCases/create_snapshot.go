package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"diffme.dev/diffme-api/internal/modules/snapshots/services"
	"log"
)

func (u *SnapshotUseCases) CreateSnapshot(snapshotParams domain.CreateSnapshotParams) (domain.Snapshot, error) {

	// Note: the ID on this is actually what we consider a reference id
	lastSnapshot, err := u.snapshotRepo.FindMostRecentByReference(snapshotParams.ID)

	snapshot, err := u.snapshotRepo.Create(snapshotParams)

	log.Printf("last snapshot: %+v", lastSnapshot)
	log.Printf("new snapshot: %+v", snapshot)

	services.SnapshotCreated(u.taskserver, lastSnapshot, snapshot)

	return snapshot, err
}
