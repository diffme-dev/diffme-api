package UseCases

import domain "diffme.dev/diffme-api/modules/snapshots"

type getSnapshotById struct {
}

func (c *snapshotUseCases) GetSnapshotByID(id string) (domain.Snapshot, error) {
	return domain.Snapshot{}, nil
}
