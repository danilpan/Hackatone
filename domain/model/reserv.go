package model

import "time"

type NewReserv struct {
	TableId  int       `json:"table_id" db:"table_id"`
	UserIIN  string    `json:"user_iin" db:"user_iin"`
	Persons  int       `json:"persons" db:"persons"`
	TimeFrom time.Time `json:"time_from" db:"time_from"`
	TimeTo   time.Time `json:"time_to" db:"time_to"`
}

type ReservDo struct {
	ID int `json:"id"`
}
