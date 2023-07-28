package services

import (
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/internal/repository"
)

type Service struct {
	Restoran Restorans
}

func New(cfg configs.Configs, repo *repository.Repository) *Service {
	return &Service{
		Restoran: NewRestoran(cfg, repo),
	}
}
