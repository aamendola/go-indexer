package elasticsearch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	logutils.Panic(err)

	m := messageUpdate{
		Doc: doc{
			Content: content,
		},
	}

	b, err := json.Marshal(m)
	logutils.Panic(err)

	update, err := client.Update(index, id, strings.NewReader(string(b)))
	logutils.Panic(err)

	log.Println(update)
	logutils.Info(id, update)

	return nil
}

// Update ...
func (e Elasticsearch) Update2(index, id string, message interface{}) error {

	cfg := elasticsearch.Config{
		Addresses: []string{
			e.uri,
		},
		Transport: &http.Transport{Proxy: nil},
	}

	client, err := elasticsearch.NewClient(cfg)
	logutils.Panic(err)

	b, err := json.Marshal(message)
	logutils.Panic(err)

	update, err := client.Update(index, id, strings.NewReader(string(b)))
	logutils.Panic(err)

	log.Println(update)
	logutils.Info(id, update)

	return nil
}