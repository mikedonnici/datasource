# datasource
Convenience package for organising multiple data source connections.

## Overview

This package allows you to create a bunch of data source connections and hang them 
off your server for much convenience.

Currently, includes connectors for:

- MongoDB
- PostgreSQL
- Pulsar
- SOLR

Todo:

- Redis
- MySQL

## Usage

A quick example:

```go
// A service...
srvc := struct {
	Connections datasource.Connections
}

// Add connections...
srvc.Connections.AddPostgresConnection("pg1", "postgres://.....")
srvc.Connections.AddPostgresConnection("pg2", "postgres://.....")
srvc.Connections.AddMongoConnection("mg1", "mongodb://.....")
```
