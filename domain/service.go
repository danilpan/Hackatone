package domain

import (
	"context"
	"log"

	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain/model"
	"github.com/madxiii/hackatone/domain/storage/db"
)

type Service interface {
	GetEstablishmentTypes(context.Context) ([]model.EstablishmentType, error)
	GetEstablishments(context.Context) ([]model.Establishment, error)
	GetEstablishment(ctx context.Context, id int) (model.Establishment, error)
	Reserv(ctx context.Context, body model.NewReserv) error
	Approve(ctx context.Context, body model.ReservDo) error
	Decline(ctx context.Context, body model.ReservDo) error
}

type service struct {
	cfg  *configs.Configs
	lgr  *log.Logger
	repo db.Repo
}

func NewService(cfg *configs.Configs, lgr *log.Logger, repo db.Repo) Service {
	return service{
		cfg:  cfg,
		lgr:  lgr,
		repo: repo,
	}
}

func (s service) GetEstablishmentTypes(ctx context.Context) ([]model.EstablishmentType, error) {
	ets, errGet := s.repo.GetEstablishmentTypes(ctx)
	if errGet != nil {
		return nil, errGet
	}

	establishmentTypes := make([]model.EstablishmentType, 0, len(ets))
	for _, et := range ets {
		establishmentTypes = append(establishmentTypes, model.EstablishmentType{
			ID:   et.ID,
			Name: et.Name,
		})
	}

	return establishmentTypes, nil
}

func (s service) GetEstablishments(ctx context.Context) ([]model.Establishment, error) {
	es, errGet := s.repo.GetEstablishments(ctx)
	if errGet != nil {
		return nil, errGet
	}

	establishments := make([]model.Establishment, 0, len(es))
	for _, e := range es {
		establishments = append(establishments, model.Establishment{
			Name:         e.Name,
			Address:      e.Address,
			Type:         e.TypeName,
			AverageCheck: e.AverageCheck,
			Rating:       e.Rating,
			ImagesURLs:   e.ImagesURLs,
		})
	}

	return establishments, nil
}

func (s service) GetEstablishment(ctx context.Context, id int) (model.Establishment, error) {
	dbEstablishment, dbTables, errGet := s.repo.GetEstablishment(ctx, id)
	if errGet != nil {
		return model.Establishment{}, errGet
	}

	tables := make([]model.Table, 0, len(dbTables))
	for _, dbTable := range dbTables {
		tables = append(tables, model.Table{
			Number:  dbTable.Number,
			Persons: dbTable.Persons,
		})
	}

	return model.Establishment{
		Name:         dbEstablishment.Name,
		Address:      dbEstablishment.Address,
		Type:         dbEstablishment.TypeName,
		AverageCheck: dbEstablishment.AverageCheck,
		Rating:       dbEstablishment.Rating,
		ImagesURLs:   dbEstablishment.ImagesURLs,
		Tables:       tables,
	}, nil
}

func (s service) Reserv(ctx context.Context, body model.NewReserv) error {
	if err := s.repo.InsertReserv(ctx, body); err != nil {
		log.Printf("Reserv InsertReserv err: %v\n", err)
		return err
	}

	return nil
}

func (s service) Approve(ctx context.Context, body model.ReservDo) error {
	if err := s.repo.UpdReserv(ctx, body, true); err != nil {
		log.Printf("Approve UpdReserv err: %v\n", err)
		return err
	}

	return nil
}

func (s service) Decline(ctx context.Context, body model.ReservDo) error {
	if err := s.repo.UpdReserv(ctx, body, false); err != nil {
		log.Printf("Decline UpdReserv err: %v\n", err)
		return err
	}

	return nil
}
