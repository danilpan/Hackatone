package handler

import (
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/internal/services"
)

type Restoran struct {
	cfg     configs.Configs
	service *services.Service
}

func NewRestoran(c configs.Configs, s *services.Service) *Restoran {
	return &Restoran{
		cfg:     c,
		service: s,
	}
}
