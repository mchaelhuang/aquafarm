//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/mchaelhuang/aquafarm/internal/app"
)

func ProvideRESTApp() *app.RESTApp {
	wire.Build(RESTAppSet)

	// Return any struct that exist inside the build
	return &app.RESTApp{}
}
