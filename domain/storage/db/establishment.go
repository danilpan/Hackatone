package db

import (
	"github.com/lib/pq"
)

type Establishment struct {
	ID           int            `db:"id"`
	Name         string         `db:"name"`
	Address      string         `db:"address"`
	TypeID       int            `db:"type_id"`
	TypeName     string         `db:"type_name"`
	AverageCheck int            `db:"average_check"`
	Rating       int            `db:"rating"`
	ImagesURLs   pq.StringArray `db:"images_urls"`
}
