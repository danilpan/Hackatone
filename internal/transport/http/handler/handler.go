package handler

import (
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/internal/services"
)

type Handler struct {
	Restoran *Restoran
}

func New(c configs.Configs, s *services.Service) *Handler {
	return &Handler{
		Restoran: NewRestoran(c, s),
	}
}
