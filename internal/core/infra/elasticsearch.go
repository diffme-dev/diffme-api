package infra

import "github.com/elastic/go-elasticsearch"

var cfg = elasticsearch.Config{
	Addresses: []string{
		"http://localhost:9200",
		"http://localhost:9201",
	},
}

func NewElasticSearch() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
