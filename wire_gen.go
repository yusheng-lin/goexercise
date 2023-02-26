// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"goexercise/api"
	"goexercise/config"
	"goexercise/db"
	"goexercise/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initapp() (*Server, func(), error) {
	config, err := ConfigProvider()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := GormProvider(config)
	if err != nil {
		return nil, nil, err
	}
	iUserRepository := db.NewUserRepository(gormDB)
	userService := service.NewUserService(iUserRepository)
	userController := api.NewUserController(userService)
	accountService := service.NewAccountService(iUserRepository)
	jwtService := service.NewJWTService(config)
	accountController := api.NewAccountController(accountService, jwtService)
	server := NewServer(userController, accountController, jwtService)
	return server, func() {
	}, nil
}

// wire.go:

var cf *config.Config

var orm *gorm.DB

func ConfigProvider() (*config.Config, error) {
	if cf == nil {
		c, err := config.NewConfig("./app.env")

		if err != nil {
			return c, err
		}
		cf = c
	}
	return cf, nil
}

func GormProvider(cf2 *config.Config) (*gorm.DB, error) {
	if orm == nil {
		o, err := gorm.Open(postgres.Open(cf2.DB_CONNSTR), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		orm = o
	}
	return orm, nil
}

var providerSet = wire.NewSet(
	ConfigProvider,
	GormProvider, db.NewUserRepository, service.NewUserService, service.NewAccountService, service.NewJWTService, api.NewUserController, api.NewAccountController, NewServer,
)