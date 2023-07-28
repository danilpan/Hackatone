package app

import (
	"github.com/madxiii/hackatone/configs"
)

type App interface {
	Run() error
}

type app struct {
	cfg configs.Configs
}

func New() (App, error) {
	cfg, err := configs.New()
	if err != nil {
		return nil, err
	}

	return app{cfg: cfg}, nil
}

func (a app) Run() error {
	return nil
}
