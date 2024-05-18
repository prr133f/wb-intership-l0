package broker

import (
	"l0/internal/database"
	"sync"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Broker struct {
	Log  *zap.Logger
	NATS *nats.Conn
	DB   *database.Database
}

var (
	natsInstance *Broker
	natsOnce     sync.Once
)

func NewBroker(DSN string, log *zap.Logger, db *database.Database) *Broker {
	natsOnce.Do(func() {
		conn, err := nats.Connect(DSN)
		if err != nil {
			log.Fatal("Unable to connect to NATS", zap.Error(err), zap.String("dsn", DSN))
		}
		natsInstance = &Broker{
			Log:  log,
			NATS: conn,
			DB:   db,
		}
	})
	return natsInstance
}
