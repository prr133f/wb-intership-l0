package domain

import (
	"encoding/json"
	"l0/internal/database"
	"l0/pkg/cache"

	"go.uber.org/zap"
)

type Domain struct {
	Log      *zap.Logger
	Database database.IFace
	Cache    cache.IFace
}

func NewDomain(log *zap.Logger, db database.IFace, cache cache.IFace) *Domain {
	return &Domain{
		Log:      log,
		Database: db,
		Cache:    cache,
	}
}

type IFace interface {
	GetDataByIDFromCache(orderUID string) (json.RawMessage, error)
}
