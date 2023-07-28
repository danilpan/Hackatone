package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/madxiii/hackatone/configs"
	"github.com/madxiii/hackatone/domain"
)

type Handler struct {
	cfg *configs.Configs
	lgr *log.Logger
	svc domain.Service
}

func New(cfg *configs.Configs, lgr *log.Logger, svc domain.Service) Handler {
	return Handler{
		cfg: cfg,
		lgr: lgr,
		svc: svc,
	}
}

func (h Handler) GetEstablishmentTypes(c echo.Context) error {
	ctx := c.Request().Context()

	estTypes, errGet := h.svc.GetEstablishmentTypes(ctx)
	if errGet != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errGet.Error())
	}

	return c.JSON(http.StatusOK, estTypes)
}

func (h Handler) GetEstablishments(c echo.Context) error {
	ctx := c.Request().Context()

	ests, errGet := h.svc.GetEstablishments(ctx)
	if errGet != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errGet.Error())
	}

	return c.JSON(http.StatusOK, ests)
}

func (h Handler) GetEstablishment(c echo.Context) error {
	ctx := c.Request().Context()

	rq := getEstablishmentRq{}
	if errBind := c.Bind(&rq); errBind != nil {
		return errBind
	}

	est, errGet := h.svc.GetEstablishment(ctx, rq.ID)
	if errGet != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errGet.Error())
	}

	return c.JSON(http.StatusOK, est)
}
