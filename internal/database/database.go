package database

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Database struct {
	Log *zap.Logger
	PG  *Postgres
}

func NewDatabase(log *zap.Logger, db *Postgres) IFace {
	return &Database{
		Log: log,
		PG:  db,
	}
}

type IFace interface {
	SetData(orderUID string, data json.RawMessage) error
	GetAllData() ([]json.RawMessage, error)
}
