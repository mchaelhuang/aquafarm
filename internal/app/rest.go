package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/delivery/middleware"
)

type RESTApp struct {
	cfg           *config.Cfg
	logger        *zap.Logger
	echo          *echo.Echo
	farmDelivery  internal.FarmDelivery
	pondDelivery  internal.PondDelivery
	statsDelivery internal.StatsDelivery
	middleware    *middleware.Middleware
}

func NewRESTApp(
	cfg *config.Cfg,
	logger *zap.Logger,
	farmDelivery internal.FarmDelivery,
	pondDelivery internal.PondDelivery,
	statsDelivery internal.StatsDelivery,
	middleware *middleware.Middleware,
) *RESTApp {
	return &RESTApp{
		cfg:           cfg,
		logger:        logger,
		echo:          echo.New(),
		farmDelivery:  farmDelivery,
		pondDelivery:  pondDelivery,
		statsDelivery: statsDelivery,
		middleware:    middleware,
	}
}

const (
	_ShutdownTimeout = 10 * time.Second
)

func (r *RESTApp) Run() error {
	r.logger.Info("Starting aquafarm-rest app")

	// Middleware
	r.echo.Use(r.middleware.StatsCounterMiddleware)

	// Route
	r.registerRoute()

	// Start server
	err := r.echo.Start(r.cfg.App.Address)
	if errors.Is(err, http.ErrServerClosed) {
		// Ignore shutdown call error
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *RESTApp) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), _ShutdownTimeout)
	defer cancel()
	if err := r.echo.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
