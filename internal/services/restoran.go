package services

import (
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/internal/repository"
)

type Restorans interface {
}

type Restoran struct {
	cfg configs.Configs
	r   *repository.Repository
}

func NewRestoran(cfg configs.Configs, repo *repository.Repository) Restorans {
	return &Restoran{
		cfg: cfg,
		r:   repo,
	}
}
