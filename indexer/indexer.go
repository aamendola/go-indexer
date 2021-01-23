package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	utils "github.com/aamendola/go-utils"
	"github.com/elastic/go-elasticsearch/v6"
)

// Elastic ...
type Elastic struct {
	index string
	uri   string
}

// Document ...
type Document struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	SystemOwner string `json:"systemowner"`
	UserOwner   string `json:"userowner"`
	Path        string `json:"path"`
}

type messageUpdate struct {
	Doc doc `json:"doc"`
}

type doc struct {
	Content string `json:"content"`
}

// NewElastic ...
func NewElastic(host, port, index string) (*Elastic, error) {

	if len(host) == 0 || len(port) == 0 || len(index) == 0 {
		return nil, fmt.Errorf("Eslastic configuration error")
	}

	elastic := Elastic{
		index: index,
		uri:   fmt.Sprintf("http://%v:%v", host, port),
	}

	return &elastic, nil
}

// UpdateDoc ...
func (e Elastic) UpdateDoc(document Document) (string, error) {

	ctx := context.Background()
	_ = ctx

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
			Content: document.Content,
		},
	}

	b, err := json.Marshal(m)
	utils.PanicOnError(err)

	update, err := client.Update(e.index, document.ID, strings.NewReader(string(b)))
	utils.PanicOnError(err)

	log.Println(update)
	return "", nil
}
