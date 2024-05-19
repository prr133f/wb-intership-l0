package database

import "go.uber.org/zap"

type Database struct {
	Log *zap.Logger
	PG  *Postgres
}

func NewDatabase(log *zap.Logger, db *Postgres) *Database {
	return &Database{
		Log: log,
		PG:  db,
	}
}
