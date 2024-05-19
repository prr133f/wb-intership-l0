package database

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (db *Database) SetData(orderUID string, data json.RawMessage) error {
	if _, err := db.PG.DB.Query(context.Background(), `
	INSERT INTO data_schema.data(order_uid, data)
	VALUES ($1, $2)`, orderUID, data); err != nil {
		db.Log.Error("Unable to insert data", zap.Error(err))
		return errors.WithStack(err)
	}

	return nil
}

func (db *Database) GetData(orderUID string) (json.RawMessage, error) {
	var data json.RawMessage
	if err := db.PG.DB.QueryRow(context.Background(), `
	SELECT data 
	FROM data_schema.data 
	WHERE order_uid = $1`, orderUID).Scan(&data); err != nil {
		db.Log.Error("Unable to get data", zap.Error(err))
		return nil, errors.WithStack(err)
	}

	return data, nil
}
