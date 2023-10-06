package db

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	logger.Debug(fmt.Sprintf("%s:%d", host, port))
	Client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d", host, port))
	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}

