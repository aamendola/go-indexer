package elasticsearch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	utils "github.com/aamendola/go-utils"
	logutils "github.com/aamendola/go-utils/log"
	"github.com/elastic/go-elasticsearch/v6"
)

// Elasticsearch ...
type Elasticsearch struct {
	uri string
}

type messageUpdate struct {
	Doc doc `json:"doc"`
}

type doc struct {
	Content string `json:"content"`
}

// NewElasticsearch ...
func NewElasticsearch(host, port, index string) (Elasticsearch, error) {

	if len(host) == 0 || len(port) == 0 || len(index) == 0 {
		return Elasticsearch{}, fmt.Errorf("Eslastic configuration error")
	}

	elastic := Elasticsearch{
		uri: fmt.Sprintf("http://%v:%v", host, port),
	}

	return elastic, nil
}

// Update ...
func (e Elasticsearch) Update(index, id, content string) error {

	cfg := elasticsearch.Config{
		Addresses: []string{
			e.uri,
		},
		Transport: &http.Transport{Proxy: nil},
	}

	client, err := elasticsearch.NewClient(cfg)
	utils.PanicOnError(err)

	m := messageUpdate{
		Doc: doc{
			Content: content,
		},
	}

	b, err := json.Marshal(m)
	utils.PanicOnError(err)

	update, err := client.Update(index, id, strings.NewReader(string(b)))
	utils.PanicOnError(err)

	log.Println(update)
	logutils.Info(id, update)

	return nil
}
