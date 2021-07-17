// Package datasource provides a structure within which to store connections to multiple data sources and services.
package datasource

import (
	"fmt"
	"log"
	"time"
)

// Connections represents connections to various platform services, most likely databases.
type Connections struct {
	Mongo    map[string]*mongoConnection
	Postgres map[string]*postgresConnection
	Pulsar   map[string]*pulsarConnection
	Solr     map[string]*solrConnection
}

// New returns a pointer to a Connections value no connections attached.
func New() *Connections {
	return &Connections{
		Mongo:    make(map[string]*mongoConnection),
		Postgres: make(map[string]*postgresConnection),
		Pulsar:   make(map[string]*pulsarConnection),
		Solr:     make(map[string]*solrConnection),
	}
}

// AddMongoConnection adds a connection to a mongo database identified by the specified key.
func (mc *Connections) AddMongoConnection(key, dsn, db string) error {
	log.Printf("Adding MONGO connection: key = '%s', db = '%s'", key, db)
	c, err := connectMongo(dsn, db)
	if err != nil {
		return fmt.Errorf("could not connect to mongo, err = %w", err)
	}
	mc.Mongo[key] = c
	return nil
}

// MongoConnByKey returns the mongodb connection value at the specified key.
func (mc *Connections) MongoConnByKey(key string) (*mongoConnection, error) {
	c, ok := mc.Mongo[key]
	if !ok {
		return nil, fmt.Errorf("no mongodb connection with key = %s", key)
	}
	return c, nil
}

// OnlyMongoConnection is a convenience function that returns the Mongo connection if there is only one.
func (mc *Connections) OnlyMongoConnection() (*mongoConnection, error) {
	num := len(mc.Mongo)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique mongo connection as %d exist", num)
	}
	var key string
	for k := range mc.Mongo { // there's only one
		key = k
	}
	return mc.Mongo[key], nil
}

// AddPostgresConnection adds a connection to a potsgres database identified by the specified key.
func (mc *Connections) AddPostgresConnection(key, dsn string) error {
	log.Printf("Adding POSTRES connection: key = '%s'", key)
	c, err := connectPostgres(dsn)
	if err != nil {
		return fmt.Errorf("could not connect to postgres, err = %w", err)
	}
	mc.Postgres[key] = c
	return nil
}

// PostgresConnByKey returns the postgres connection value at the specified key.
func (mc *Connections) PostgresConnByKey(key string) (*postgresConnection, error) {
	c, ok := mc.Postgres[key]
	if !ok {
		return nil, fmt.Errorf("no postgresdb connection with key = %s", key)
	}
	return c, nil
}

// OnlyPostgresConnection is a convenience function that returns the Postgres connection if there is only one.
func (mc *Connections) OnlyPostgresConnection() (*postgresConnection, error) {
	num := len(mc.Postgres)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique postgres connection as %d exist", num)
	}
	var key string
	for k := range mc.Postgres {
		key = k
	}
	return mc.Postgres[key], nil
}

// AddPulsarConnection adds a connection to a Pulsar server identified by the specified key.
func (mc *Connections) AddPulsarConnection(key, dsn string, timeoutSeconds time.Duration) error {
	log.Printf("Adding PULSAR connection: key = '%s'", key)
	c, err := connectPulsar(dsn, timeoutSeconds)
	if err != nil {
		return fmt.Errorf("could not connect to pulsar, err = %w", err)
	}
	mc.Pulsar[key] = c
	return nil
}

// PulsarConnByKey returns the Pulsar connection value at the specified key.
func (mc *Connections) PulsarConnByKey(key string) (*pulsarConnection, error) {
	c, ok := mc.Pulsar[key]
	if !ok {
		return nil, fmt.Errorf("no Pulsar connection with key = %s", key)
	}
	return c, nil
}

// OnlyPulsarConnection is a convenience function that returns the Pulsar connection if there is only one.
func (mc *Connections) OnlyPulsarConnection() (*pulsarConnection, error) {
	num := len(mc.Pulsar)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique Pulsar connection as %d exist", num)
	}
	var key string
	for k := range mc.Pulsar {
		key = k
	}
	return mc.Pulsar[key], nil
}

// AddSolrConnection adds a connection to a Solr server identified by the specified key.
func (mc *Connections) AddSolrConnection(key, dsn, core string, timeoutSeconds time.Duration) error {
	log.Printf("Adding SOLR connection: key = '%s'", key)
	c, err := connectSolr(dsn, core, timeoutSeconds)
	if err != nil {
		return fmt.Errorf("could not connect to solr, err = %w", err)
	}
	mc.Solr[key] = c
	return nil
}

// SolrConnByKey returns the Solr connection value at the specified key.
func (mc *Connections) SolrConnByKey(key string) (*solrConnection, error) {
	c, ok := mc.Solr[key]
	if !ok {
		return nil, fmt.Errorf("no Solr connection with key = %s", key)
	}
	return c, nil
}

// OnlySolrConnection is a convenience function that returns the Solr connection if there is only one.
func (mc *Connections) OnlySolrConnection() (*solrConnection, error) {
	num := len(mc.Solr)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique Solr connection as %d exist", num)
	}
	var key string
	for k := range mc.Solr {
		key = k
	}
	return mc.Solr[key], nil
}
