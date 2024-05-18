package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	Log *zap.Logger
	DB  *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPostgres(ctx context.Context, DSN string, log *zap.Logger) *Postgres {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, DSN)
		if err != nil {
			log.Fatal("Unable to connect to database", zap.Error(err))
		}

		pgInstance = &Postgres{
			Log: log,
			DB:  db,
		}
	})
	return pgInstance
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}
