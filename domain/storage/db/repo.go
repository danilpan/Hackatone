package db

import (
	"context"
)

type Repo interface {
	GetEstablishmentTypes(context.Context) ([]EstablishmentType, error)
	GetEstablishments(context.Context) ([]Establishment, error)
}
