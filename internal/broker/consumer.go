package broker

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Listen listens for messages on the NATS JetStream and saves them to the database and cache.
//
// It returns an error if there was a problem creating the JetStream context or subscribing to the JetStream.
func (b *Broker) Listen() error {
	js, err := b.NATS.JetStream()
	if err != nil {
		b.Log.Error("Unable to create JetStream context", zap.Error(err))
		return errors.WithStack(err)
	}

	if _, err = js.Subscribe("l0.*", func(m *nats.Msg) {
		b.Log.Info("Message received", zap.String("subject", m.Subject), zap.ByteString("data", m.Data))

		uid, subErr := b.SaveToDB(m) // TODO: Отлошить ошибку pkey и не записывать в кэш
		if subErr != nil {
			b.Log.Error("Unable to save data", zap.Error(err))
			return
		}

		b.Cache.Set(uid, json.RawMessage(m.Data))
	}, nats.Durable("app")); err != nil {
		b.Log.Error("Unable to subscribe to JetStream", zap.Error(err))
		return errors.WithStack(err)
	}

	return nil
}
