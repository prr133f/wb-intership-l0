package database

import "go.uber.org/zap"

type Database struct {
	Log *zap.Logger
	DB  *Postgres
}

func NewDatabase(log *zap.Logger, db *Postgres) *Database {
	return &Database{
		Log: log,
		DB:  db,
	}
}
