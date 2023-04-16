package internal

import (
	"github.com/labstack/echo/v4"
)

type FarmDelivery interface {
	GetList(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type PondDelivery interface {
	GetList(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type StatsDelivery interface {
	Get(c echo.Context) error
}
