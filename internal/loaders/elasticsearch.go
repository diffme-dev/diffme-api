package loaders

import (
	"github.com/elastic/go-elasticsearch"
	"log"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}