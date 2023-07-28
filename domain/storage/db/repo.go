package db

import (
	"context"
)

type Repo interface {
	GetEstablishmentTypes(context.Context) ([]EstablishmentType, error)
	GetEstablishments(context.Context) ([]Establishment, error)
	GetEstablishment(ctx context.Context, id int) (Establishment, []Table, error)
}
