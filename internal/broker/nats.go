package broker

import (
	"l0/internal/database"
	"l0/pkg/cache"
	"sync"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Broker struct {
	Log   *zap.Logger
	NATS  *nats.Conn
	DB    database.IFace
	Cache cache.IFace
}

type IFace interface {
	Listen() error
	SaveToDB(msg *nats.Msg) (string, error)
}

func NewBroker(dsn string, log *zap.Logger, db database.IFace, cache cache.IFace) IFace {
	var (
		natsInstance *Broker
		natsOnce     sync.Once
	)

	natsOnce.Do(func() {
		conn, err := nats.Connect(dsn)
		if err != nil {
			log.Fatal("Unable to connect to NATS", zap.Error(err), zap.String("dsn", dsn))
		}
		natsInstance = &Broker{
			Log:   log,
			NATS:  conn,
			DB:    db,
			Cache: cache,
		}
	})
	return natsInstance
}
