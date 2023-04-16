package delivery

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mchaelhuang/aquafarm/internal/constant"
)

func handlerErrorResponse(err error) error {
	if errors.Is(err, constant.ErrNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if errors.Is(err, constant.ErrNotImplemented) {
		return echo.NewHTTPError(http.StatusNotImplemented, err.Error())
	}
	if errors.Is(err, constant.ErrIncorrectFarmID) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
