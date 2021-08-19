package infra

import (
	"diffme.dev/diffme-api/config"
	"github.com/elastic/go-elasticsearch"
)

var cfg = elasticsearch.Config{
	Addresses: []string{
		config.GetConfig().ElasticUri,
	},
}

func NewElasticSearch() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
