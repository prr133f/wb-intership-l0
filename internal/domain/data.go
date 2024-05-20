package domain

import (
	"encoding/json"
	"errors"

	"go.uber.org/zap"
)

func (d *Domain) GetDataByIDFromCache(orderUID string) (json.RawMessage, error) {
	data, ok := d.Cache.Get(orderUID)

	if !ok {
		d.Log.Error("No such data in cache", zap.String("order_uid", orderUID))
		return nil, errors.New("no such data in cache")
	}

	return data.(json.RawMessage), nil
}
