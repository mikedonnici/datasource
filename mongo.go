package datasource

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoConnection struct {
	dsn string
	*mongo.Client
	*mongo.Database
}

// newMongoConnection returns a new MongoDB connection which includes the client and a database handle.
func newMongoConnection(dsn, dbname string) (*mongoConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn := mongoConnection{
		dsn:    dsn,
		Client: nil,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, fmt.Errorf("could not create mongo client, err = %w", err)
	}
	conn.Client = client
	if err := conn.Check(); err != nil {
		return nil, fmt.Errorf("mongo dbname connection check failed, err = %w", err)
	}
	conn.Database = conn.Client.Database(dbname)

	return &conn, nil
}

// connectMongo returns a mongodb connection or an error.
func connectMongo(dsn, db string) (*mongoConnection, error) {
	return newMongoConnection(dsn, db)
}

// Check verifies the connection to the database and returns an error if there's a problem.
// Note: This is better than ping because it forces a round trip to the database.
func (c *mongoConnection) Check() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.Ping(ctx, readpref.Primary())
}
