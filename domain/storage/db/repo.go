package db

import (
	"context"

	"github.com/madxiii/hackatone/domain/model"
)

type Repo interface {
	GetEstablishmentTypes(context.Context) ([]EstablishmentType, error)
	GetEstablishments(context.Context) ([]Establishment, error)
	GetEstablishment(ctx context.Context, id int) (Establishment, []Table, error)
	InsertReserv(ctx context.Context, body model.NewReserv) error
	UpdReserv(ctx context.Context, body model.ReservDo, cofirm bool) error
}
