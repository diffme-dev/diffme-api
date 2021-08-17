package infra

import "github.com/elastic/go-elasticsearch"

var cfg = elasticsearch.Config{
	Addresses: []string{
		"http://localhost:9200",
	},
}

func NewElasticSearch() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
