package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/madxiii/hackatone/configs"
)

func InitDB(ctx context.Context, cfg configs.Configs) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, cfg.Store.DB.Driver, cfg.Store.DB.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(time.Duration(120) * time.Second)

	return db, nil
}
