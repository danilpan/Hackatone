package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Restoran Restorans
}

func New(db *sqlx.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Restoran: New,
	}
}
