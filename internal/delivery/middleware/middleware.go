package middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type Middleware struct {
	logger  *zap.Logger
	statsUC internal.StatsUC
}

func NewMiddleware(
	logger *zap.Logger,
	statsUC internal.StatsUC,
) *Middleware {
	return &Middleware{
		logger:  logger,
		statsUC: statsUC,
	}
}

func (m *Middleware) StatsCounterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		info := entity.StatsRequestInfo{
			Method:    c.Request().Method,
			Endpoint:  c.Request().RequestURI,
			UserAgent: c.Request().UserAgent(),
		}
		err := m.statsUC.Add(c.Request().Context(), info)
		if err != nil {
			return err
		}
		return next(c)
	}
}
