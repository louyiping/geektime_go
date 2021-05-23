// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/beer-shop/app/payment/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/data"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/server"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	beerRepo := data.NewBeerRepo(dataData, logger)
	beerUseCase := biz.NewBeerUseCase(beerRepo, logger)
	paymentService := service.NewPaymentService(beerUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, paymentService)
	grpcServer := server.NewGRPCServer(confServer, paymentService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}