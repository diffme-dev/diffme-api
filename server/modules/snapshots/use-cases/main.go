package use_cases

import (
	"diffme.dev/diffme-api/server/core/interfaces"
	SnapshotDomain "diffme.dev/diffme-api/server/modules/snapshots"
)

type SnapshotUseCases struct {
	snapshotRepo SnapshotDomain.SnapshotRepo
	compression  interfaces.Compression
}

func NewSnapshotUseCases(snapshotRepo SnapshotDomain.SnapshotRepo, compression interfaces.Compression) SnapshotDomain.SnapshotUseCases {
	return &SnapshotUseCases{
		snapshotRepo: snapshotRepo,
		compression:  compression,
	}
}
