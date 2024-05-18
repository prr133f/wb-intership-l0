package domain

import (
	"l0/internal/database"

	"go.uber.org/zap"
)

type Domain struct {
	Log      *zap.Logger
	Database *database.Database
}

func NewDomain(log *zap.Logger, db *database.Database) *Domain {
	return &Domain{
		Log:      log,
		Database: db,
	}
}
