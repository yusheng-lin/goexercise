// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func initapp() (*Server, func(), error) {
	config, err := configProvider()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gormProvider(config)
	if err != nil {
		return nil, nil, err
	}
	iUserRepository := db.NewUserRepository(gormDB)
	userService := service.NewUserService(iUserRepository)
	userController := api.NewUserController(userService)
	accountService := service.NewAccountService(iUserRepository)
	jwtService := service.NewJWTService(config)
	accountController := api.NewAccountController(accountService, jwtService)
	server := NewServer(userController, accountController, jwtService, config)
	return server, func() {
	}, nil
}

// wire.go:

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

func gormProvider(cf2 *config.Config) (*gorm.DB, error) {
	if orm == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Taipei", cf2.POSTGRES_HOST, cf2.POSTGRES_USER, cf2.POSTGRES_PASSWORD, cf2.POSTGRES_DB)
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
	gormProvider, db.NewUserRepository, service.NewUserService, service.NewAccountService, service.NewJWTService, api.NewUserController, api.NewAccountController, NewServer,
)
