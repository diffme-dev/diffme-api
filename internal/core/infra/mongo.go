package infra

import (
	"diffme.dev/diffme-api/internal/config"
	"github.com/go-bongo/bongo"
)

func NewBongoConnection() (*bongo.Connection, error) {
	client, err := bongo.Connect(&bongo.Config{
		ConnectionString: config.GetConfig().MongoUri,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
