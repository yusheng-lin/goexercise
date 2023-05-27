//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
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

func configProvider() (*config.Config, error) {
	if cf == nil {
		c, err := config.NewConfig("./env")

		if err != nil {
			return c, err
		}
		cf = c
	}
	return cf, nil
}

func gormProvider(cf *config.Config) (*gorm.DB, error) {
	if orm == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Taipei", cf.POSTGRES_HOST, cf.POSTGRES_USER, cf.POSTGRES_PASSWORD, cf.POSTGRES_DB)
		o, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		orm = o
	}
	return orm, nil
}

var providerSet = wire.NewSet(
	configProvider,
	gormProvider,
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
