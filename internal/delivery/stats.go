package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mchaelhuang/aquafarm/internal"
)

type StatsDelivery struct {
	statsUC internal.StatsUC
}

func NewStatsDelivery(
	statsUC internal.StatsUC,
) *StatsDelivery {
	return &StatsDelivery{
		statsUC: statsUC,
	}
}

func (s *StatsDelivery) Get(c echo.Context) error {
	result, err := s.statsUC.Get(c.Request().Context())
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, result)
}
