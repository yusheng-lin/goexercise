package main

import (
	"fmt"
	"goexercise/api"
	"goexercise/config"
	"goexercise/service"

	docs "goexercise/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router            *gin.Engine
	usercontroller    *api.UserController
	accountcontroller *api.AccountController
	jwtservice        *service.JWTService
	config            *config.Config
}

func NewServer(usercontroller *api.UserController, accountcontroller *api.AccountController, jwtservice *service.JWTService, config *config.Config) *Server {
	return &Server{
		usercontroller:    usercontroller,
		accountcontroller: accountcontroller,
		jwtservice:        jwtservice,
		router:            gin.Default(),
		config:            config,
	}
}

func (server *Server) Run(port int) {
	p := fmt.Sprintf(":%d", port)

	if server.config.RUN_TLS {
		server.router.RunTLS(p, "./certs/server.crt", "./certs/server.key")
	} else {
		server.router.Run(p)
	}
}

func (server *Server) SetupRouter() {
	server.router.POST("/api/v1/Login", server.accountcontroller.Login)
	server.router.POST("/api/v1/SignUp", server.accountcontroller.SignUp)
	server.router.POST("/api/v1/user", server.jwtservice.AuthRequired, server.usercontroller.Users)
	server.router.GET("/api/v1/user/:name", server.jwtservice.AuthRequired, server.usercontroller.User)
	server.router.DELETE("/api/v1/user", server.jwtservice.AuthRequired, server.accountcontroller.Delete)
	server.router.PUT("/api/v1/user", server.jwtservice.AuthRequired, server.accountcontroller.Update)
	docs.SwaggerInfo.BasePath = "/api/v1"
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
