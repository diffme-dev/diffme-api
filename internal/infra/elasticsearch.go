package infra

import (
	"github.com/elastic/go-elasticsearch"
	"log"
)

func NewElasticSearch() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}
