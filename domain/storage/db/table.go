package db

type Table struct {
	ID              int `db:"id"`
	EstablishmentID int `db:"establishment_id"`
	Number          int `db:"number"`
	Persons         int `db:"persons"`
}
