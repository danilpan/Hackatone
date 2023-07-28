package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Restorans interface {
}

type Restoran struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewRestoran(db *sqlx.DB, rdb *redis.Client) Restorans {
	return &Restoran{
		db:  db,
		rdb: rdb,
	}
}
