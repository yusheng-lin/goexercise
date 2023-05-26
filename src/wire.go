//go:build wireinject
// +build wireinject

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

func GormProvider(cf *config.Config) (*gorm.DB, error) {
	if orm == nil {
		o, err := gorm.Open(postgres.Open(cf.DB_CONNSTR), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		orm = o
	}
	return orm, nil
}

var providerSet = wire.NewSet(
	ConfigProvider,
	GormProvider,
	db.NewUserRepository,
	service.NewUserService,
	service.NewAccountService,
	service.NewJWTService,
	api.NewUserController,
	api.NewAccountController,
	NewServer,
)

func initapp() (*Server, func(), error) {
	panic(wire.Build(providerSet))
}
