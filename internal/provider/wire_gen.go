// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/mchaelhuang/aquafarm/internal/app"
	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/delivery"
	"github.com/mchaelhuang/aquafarm/internal/delivery/middleware"
	"github.com/mchaelhuang/aquafarm/internal/repository"
	"github.com/mchaelhuang/aquafarm/internal/usecase"
)

// Injectors from wire.go:

func ProvideRESTApp() *app.RESTApp {
	cfg := config.Get()
	zapLogger := LoggerProvider()
	gormDB := DatabaseProvider(cfg)
	farmRepo := repository.NewFarmRepo(cfg, zapLogger, gormDB)
	pondRepo := repository.NewPondRepo(cfg, zapLogger, gormDB)
	farmUC := usecase.NewFarmUC(farmRepo, pondRepo)
	farmDelivery := delivery.NewFarmDelivery(cfg, farmUC)
	pondUC := usecase.NewPondUC(zapLogger, farmRepo, pondRepo)
	pondDelivery := delivery.NewPondDelivery(cfg, pondUC)
	client := RedisProvider(cfg)
	statsRepo := repository.NewStatsRepo(client)
	statsUC := usecase.NewStatsUC(zapLogger, statsRepo)
	statsDelivery := delivery.NewStatsDelivery(statsUC)
	middlewareMiddleware := middleware.NewMiddleware(zapLogger, statsUC)
	restApp := app.NewRESTApp(cfg, zapLogger, farmDelivery, pondDelivery, statsDelivery, middlewareMiddleware)
	return restApp
}
