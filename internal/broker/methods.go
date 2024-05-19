package broker

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (b *Broker) SaveToDB(msg *nats.Msg) (string, error) {
	marshalledData := make(map[string]interface{})
	err := json.Unmarshal(msg.Data, &marshalledData)
	if err != nil {
		b.Log.Error("Unable to unmarshal message", zap.Error(err))
		return "", errors.WithStack(err)
	}

	err = b.DB.SetData(marshalledData["order_uid"].(string), msg.Data)
	if err != nil {
		b.Log.Error("Unable to save data", zap.Error(err))
		return "", errors.WithStack(err)
	}

	return marshalledData["order_uid"].(string), nil
}
