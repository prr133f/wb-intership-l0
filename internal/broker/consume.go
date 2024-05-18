package broker

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func (b *Broker) Listen() error {
	js, err := b.NATS.JetStream()
	if err != nil {
		b.Log.Error("Unable to create JetStream context", zap.Error(err))
		return err
	}

	js.Subscribe("l0.*", func(m *nats.Msg) {
		b.Log.Info("Message received", zap.String("subject", m.Subject), zap.ByteString("data", m.Data))
	}, nats.Durable("app"))

	return nil
}
