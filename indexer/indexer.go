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

type ElasticSearchConfig struct {
	Host  string
	Port  string
	Index string
}

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

//  GetElasticSearchConfig()
func NewElasticSearchConfig(host, port, index string) (*ElasticSearchConfig, error) {
	elasticSearchConfig := ElasticSearchConfig{
		Host:  host,
		Port:  port,
		Index: index,
	}

	if len(elasticSearchConfig.Host) == 0 || len(elasticSearchConfig.Port) == 0 || len(elasticSearchConfig.Index) == 0 {
		return nil, fmt.Errorf("EslasticSearchConfig error")
	}
	return &elasticSearchConfig, nil
}

func (e ElasticSearchConfig) UpdateDoc(document Document) (string, error) {

	ctx := context.Background()
	_ = ctx

	esURI := fmt.Sprintf("http://%v:%v", e.Host, e.Port)

	cfg := elasticsearch.Config{
		Addresses: []string{
			esURI,
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

	update, err := client.Update(e.Index, document.ID, strings.NewReader(string(b)))
	utils.PanicOnError(err)

	log.Println(update)
	return "", nil
}
