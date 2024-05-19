package domain

import (
	"l0/internal/database"
	"l0/pkg/cache"

	"go.uber.org/zap"
)

type Domain struct {
	Log      *zap.Logger
	Database *database.Database
	Cache    *cache.Cache
}

func NewDomain(log *zap.Logger, db *database.Database, cache *cache.Cache) *Domain {
	return &Domain{
		Log:      log,
		Database: db,
		Cache:    cache,
	}
}
