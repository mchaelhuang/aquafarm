package provider

import (
	"github.com/google/wire"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/app"
	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/delivery"
	"github.com/mchaelhuang/aquafarm/internal/delivery/middleware"
	"github.com/mchaelhuang/aquafarm/internal/repository"
	"github.com/mchaelhuang/aquafarm/internal/usecase"
)

var BaseSet = wire.NewSet(
	config.Get,
	LoggerProvider,
	DatabaseProvider,
	RedisProvider,
)

var RepositorySet = wire.NewSet(
	repository.NewFarmRepo,
	wire.Bind(new(internal.FarmRepo), new(*repository.FarmRepo)),
	repository.NewPondRepo,
	wire.Bind(new(internal.PondRepo), new(*repository.PondRepo)),
	repository.NewStatsRepo,
	wire.Bind(new(internal.StatsRepo), new(*repository.StatsRepo)),
)

var UseCaseSet = wire.NewSet(
	usecase.NewFarmUC,
	wire.Bind(new(internal.FarmUC), new(*usecase.FarmUC)),
	usecase.NewPondUC,
	wire.Bind(new(internal.PondUC), new(*usecase.PondUC)),
	usecase.NewStatsUC,
	wire.Bind(new(internal.StatsUC), new(*usecase.StatsUC)),
)

var DeliverySet = wire.NewSet(
	delivery.NewFarmDelivery,
	wire.Bind(new(internal.FarmDelivery), new(*delivery.FarmDelivery)),
	delivery.NewPondDelivery,
	wire.Bind(new(internal.PondDelivery), new(*delivery.PondDelivery)),
	delivery.NewStatsDelivery,
	wire.Bind(new(internal.StatsDelivery), new(*delivery.StatsDelivery)),
)

var RESTAppSet = wire.NewSet(
	BaseSet,
	RepositorySet,
	UseCaseSet,
	DeliverySet,

	middleware.NewMiddleware,

	app.NewRESTApp,
)
