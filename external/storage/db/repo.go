package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain/model"
	"github.com/madxiii/hackatone/domain/storage/db"
)

const (
	queryGetEstablishmentTypes = `SELECT id, name FROM establishment_types;`
	queryGetEstablishments     = `
	SELECT
		e.id, e.name, e.address,
		e.type_id, et.name AS type_name,
		e.average_check, e.rating, e.images_urls
	FROM establishments e
	LEFT JOIN establishment_types et ON et.id = e.type_id;`

	queryGetEstablishment = `
	SELECT
		e.id, e.name, e.address,
		e.type_id, et.name AS type_name,
		e.average_check, e.rating, e.images_urls
	FROM establishments e
	LEFT JOIN establishment_types et ON et.id = e.type_id
	WHERE e.id = $1;`

	queryGetEstablishmentsTables = `
	SELECT t.id, t.establishment_id, t.number, t.persons
	FROM tables t
	WHERE t.establishment_id = $1;`
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

func (r repo) GetEstablishments(ctx context.Context) (es []db.Establishment, err error) {
	defer func() {
		if err == nil {
			r.lgr.Printf("get establishments: success")
			return
		}

		r.lgr.Printf("get establishments: %v", err)
	}()

	errGet := r.db.SelectContext(ctx, &es, queryGetEstablishments)
	if errGet != nil {
		return nil, fmt.Errorf("executing query: %w", errGet)
	}

	return es, nil
}

func (r repo) GetEstablishment(ctx context.Context, id int) (est db.Establishment, ts []db.Table, err error) {
	defer func() {
		if err == nil {
			r.lgr.Printf("get establishment with id %d: success", id)
			return
		}

		r.lgr.Printf("get establishment with id %d: %v", id, err)
	}()

	errGet := r.db.GetContext(ctx, &est, queryGetEstablishment, id)
	if errGet != nil {
		return db.Establishment{}, nil, fmt.Errorf("executing query: %w", errGet)
	}

	errGetTables := r.db.SelectContext(ctx, &ts, queryGetEstablishmentsTables, id)
	if errGetTables != nil {
		return db.Establishment{}, nil, fmt.Errorf("executing query: %w", errGetTables)
	}

	return est, ts, nil
}

func (r repo) InsertReserv(ctx context.Context, body model.NewReserv) error {
	query := fmt.Sprintf(`INSERT INTO reservations (table_id, user_iin, time_from, time_to, persons)
		  VALUES (:table_id, :user_iin, :time_from, :time_to, :persons)`)
	_, err := r.db.NamedExecContext(ctx, query, body)

	return err
}

func (r repo) UpdReserv(ctx context.Context, body model.ReservDo, confirm bool) error {
	query := fmt.Sprintf(`UPDATE reservation SET confirmed=$1 WHERE id=$2`)
	_, err := r.db.ExecContext(ctx, query, confirm, body.ID)

	return err
}
