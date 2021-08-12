package UseCases

import SnapshotDomain "diffme.dev/diffme-api/modules/snapshots"

type snapshotUseCases struct {
	snapshotRepo SnapshotDomain.SnapshotRepo
}

func NewSnapshotUseCases(snapshotRepo SnapshotDomain.SnapshotRepo) SnapshotDomain.SnapshotUseCases {
	return &snapshotUseCases{
		snapshotRepo: snapshotRepo,
	}
}
