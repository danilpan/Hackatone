package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/madxiii/hackatone/api/handler"
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain"
	"github.com/madxiii/hackatone/external/storage/db"
)

type App interface {
	Run() error
}

type app struct {
	cfg configs.Configs
	hnd handler.Handler
	fwk *echo.Echo
}

func New() (App, error) {
	cfg, err := configs.New()
	if err != nil {
		return nil, err
	}

	lgr := log.New(os.Stdout, "", 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo, errNewRepo := db.NewRepo(ctx, &cfg, lgr)
	if errNewRepo != nil {
		return nil, errNewRepo
	}

	service := domain.NewService(&cfg, lgr, repo)
	hnd := handler.New(&cfg, lgr, service)

	return app{
		cfg: cfg,
		hnd: hnd,
		fwk: echo.New(),
	}, nil
}

func (a app) Run() error {
	a.fwk.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	a.Routes()

	go a.start()

	errorChan := make(chan error)
	gracefulShutdown(errorChan)

	err := <-errorChan

	return err
}

func (a app) start() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := a.fwk.Start(port); err != nil {
		log.Fatalf("incorrect server shutdown: %v\n", err)
	}
}

func gracefulShutdown(errorChan chan error) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		errorChan <- fmt.Errorf("received signal %v", sig)
	}()
}
