package UseCases

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	"github.com/RichardKnop/machinery/v1"
)

type SnapshotUseCases struct {
	snapshotRepo SnapshotDomain.SnapshotRepo
	taskserver   *machinery.Server
	compression  interfaces.Compression
}

func NewSnapshotUseCases(snapshotRepo SnapshotDomain.SnapshotRepo, taskserver *machinery.Server, compression interfaces.Compression) SnapshotDomain.SnapshotUseCases {
	return &SnapshotUseCases{
		snapshotRepo: snapshotRepo,
		taskserver:   taskserver,
		compression:  compression,
	}
}
