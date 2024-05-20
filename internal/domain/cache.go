package domain

import (
	"encoding/json"
	"l0/internal/database"
	"l0/pkg/cache"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func RestoreCache(db database.IFace, cache cache.IFace, log *zap.Logger) error {
	data, err := db.GetAllData()
	if err != nil {
		log.Error("Unable to get data from database", zap.Error(err))
		return errors.WithStack(err)
	}

	for _, row := range data {
		domainData := make(map[string]interface{})
		err = json.Unmarshal(row, &domainData)
		if err != nil {
			log.Error("Unable to unmarshal data", zap.Error(err))
			return errors.WithStack(err)
		}

		cache.Set(domainData["order_uid"].(string), row)
	}

	return nil
}
