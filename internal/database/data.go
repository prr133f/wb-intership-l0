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

func (db *Database) GetAllData() ([]json.RawMessage, error) {
	var data []json.RawMessage
	rows, err := db.PG.DB.Query(context.Background(), `
	SELECT data 
	FROM data_schema.data`)
	if err != nil {
		db.Log.Error("Unable to get data", zap.Error(err))
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		var d json.RawMessage
		if err = rows.Scan(&d); err != nil {
			db.Log.Error("Unable to scan data", zap.Error(err))
			return nil, errors.WithStack(err)
		}

		data = append(data, d)
	}

	return data, nil
}
