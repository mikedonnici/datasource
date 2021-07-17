package datasource

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sf9v/solr-go"
)

type solrConnection struct {
	uri    string
	core   string
	Client solr.Client
}

// newSolrConnection
func newSolrConnection(dsn, core string, timeoutSeconds time.Duration) (*solrConnection, error) {

	r := solr.NewDefaultRequestSender().WithHTTPClient(&http.Client{Timeout: timeoutSeconds * time.Second})
	client := solr.NewJSONClient(dsn).WithRequestSender(r)

	conn := solrConnection{
		uri:    dsn,
		core:   core,
		Client: client,
	}
	if err := conn.Check(); err != nil {
		return nil, fmt.Errorf("solr connection check failed, err = %w", err)
	}

	return &conn, nil
}

// connectSolr returns a Solr connection or an error.
func connectSolr(dsn, core string, timeoutSeconds time.Duration) (*solrConnection, error) {
	return newSolrConnection(dsn, core, timeoutSeconds)
}

// Check the client connection by checking the core status
func (c *solrConnection) Check() error {
	ctx := context.Background()
	_, err := c.Client.CoreStatus(ctx, solr.NewCoreParams(c.core))
	return err
}

// Close the connection
func (c *solrConnection) Close() error {
	ctx := context.Background()
	return c.Client.UnloadCore(ctx, solr.NewCoreParams(c.core))
}
