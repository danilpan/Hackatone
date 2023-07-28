package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain/storage/db"
)

const (
	queryGetEstablishmentTypes = `SELECT id, name FROM establishment_types;`
)

type repo struct {
	cfg *configs.Configs
	lgr *log.Logger
	db  *sqlx.DB
}

func NewRepo(ctx context.Context, cfg *configs.Configs, lgr *log.Logger) (db.Repo, error) {
	database, errNewDB := newDB(ctx, cfg)
	if errNewDB != nil {
		return nil, errNewDB
	}

	return repo{
		cfg: cfg,
		lgr: lgr,
		db:  database,
	}, nil
}

func (r repo) GetEstablishmentTypes(ctx context.Context) (ets []db.EstablishmentType, err error) {
	defer func() {
		if err == nil {
			r.lgr.Printf("get establishment types: success")
			return
		}

		r.lgr.Printf("get establishment types: %v", err)
	}()

	errGet := r.db.SelectContext(ctx, &ets, queryGetEstablishmentTypes)
	if errGet != nil {
		return nil, fmt.Errorf("executing query: %w", errGet)
	}

	return ets, nil
}
