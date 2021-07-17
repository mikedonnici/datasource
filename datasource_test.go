package datasource_test

import (
	"testing"

	"github.com/mikedonnici/datasource"
)

// Each of these platform services should be represented in the
// platform docker-compose file so that the tests can be run.
const (
	mongoDSN1    = "mongodb://localhost:27018"
	mongoDSN2    = "mongodb://localhost:27019"
	postgresDSN1 = "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable"
	postgresDSN2 = "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable"
	pulsarDSN1   = "pulsar://localhost:6651"
	pulsarDSN2   = "pulsar://localhost:6652"

	// Just one Solr for now - cores make it complicated
	solrDSN1  = "http://localhost:8983"
	solrCore1 = "somecore"

	// for all, where relevant
	timeoutSecs = 5
)

func TestPlatform(t *testing.T) {

	t.Run("tests", func(t *testing.T) {
		t.Run("testAddMongoConnection", testAddMongoConnection)
		t.Run("testAddPostgresConnection", testAddPostgresConnection)
		t.Run("testAddPulsarConnection", testAddPulsarConnection)
		t.Run("testAddSolrConnection", testAddSolrConnection)
	})
}

func testAddMongoConnection(t *testing.T) {

	cases := []struct {
		dsn string
		key string
		db  string
	}{
		{mongoDSN1, "mongo-1", "db-1"}, // first conn to first server
		{mongoDSN1, "mongo-2", "db-2"}, // second conn to first server
		{mongoDSN2, "mongo-3", "db-3"}, // first conn to second server
		{mongoDSN2, "mongo-4", "db-4"}, // second conn to second server
	}

	conns := datasource.New()
	for _, c := range cases {
		if err := conns.AddMongoConnection(c.key, c.dsn, c.db); err != nil {
			t.Fatalf(".AddMongoConnection(%s, %s, %s) err = %s", c.key, c.dsn, c.db, err)
		}
	}

	// Check number of connections
	want := 4
	got := len(conns.Mongo)
	if got != want {
		t.Errorf("Number of connections = %d, want %d", got, want)
	}
}

func testAddPostgresConnection(t *testing.T) {

	cases := []struct {
		dsn string
		key string
	}{
		{postgresDSN1, "postgres-1"}, // first conn to first server
		{postgresDSN1, "postgres-2"}, // second conn to first server
		{postgresDSN2, "postgres-3"}, // first conn to second server
		{postgresDSN2, "postgres-4"}, // second conn to second server
	}

	conns := datasource.New()
	for _, c := range cases {
		if err := conns.AddPostgresConnection(c.key, c.dsn); err != nil {
			t.Fatalf(".AddPostgresConnection(%s, %s) err = %s", c.key, c.dsn, err)
		}
	}

	// Check number of connections
	want := 4
	got := len(conns.Postgres)
	if got != want {
		t.Errorf("Number of connections = %d, want %d", got, want)
	}
}

func testAddPulsarConnection(t *testing.T) {

	cases := []struct {
		dsn string
		key string
	}{
		{pulsarDSN1, "pulsar-1"}, // first conn to first server
		{pulsarDSN1, "pulsar-2"}, // second conn to first server
		{pulsarDSN2, "pulsar-3"}, // first conn to second server
		{pulsarDSN2, "pulsar-4"}, // second conn to second server
	}

	conns := datasource.New()
	for _, c := range cases {
		if err := conns.AddPulsarConnection(c.key, c.dsn, timeoutSecs); err != nil {
			t.Fatalf(".AddPulsarConnection(%s, %s) err = %s", c.key, c.dsn, err)
		}
	}

	// Check number of connections
	want := 4
	got := len(conns.Pulsar)
	if got != want {
		t.Errorf("Number of connections = %d, want %d", got, want)
	}
}

func testAddSolrConnection(t *testing.T) {

	cases := []struct {
		dsn  string
		core string
		key  string
	}{
		{solrDSN1, solrCore1, "solr-1"}, // first conn to first server
	}

	conns := datasource.New()
	for _, c := range cases {
		if err := conns.AddSolrConnection(c.key, c.dsn, c.core, timeoutSecs); err != nil {
			t.Fatalf(".AddSolrConnection(%s, %s) err = %s", c.key, c.dsn, err)
		}
	}

	// Check number of connections
	want := 1
	got := len(conns.Solr)
	if got != want {
		t.Errorf("Number of connections = %d, want %d", got, want)
	}
}
