package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/madxiii/hackatone/configs"

	_ "github.com/lib/pq"
)

func newDB(ctx context.Context, cfg *configs.Configs) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, cfg.Store.DB.Driver, cfg.Store.DB.DSN)
	if err != nil {
		return nil, err
	}

	if errPing := db.Ping(); errPing != nil {
		return nil, errPing
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(time.Duration(120) * time.Second)

	return db, nil
}
