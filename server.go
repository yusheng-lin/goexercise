package main

import (
	"fmt"
	"goexercise/api"
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
}

func NewServer(usercontroller *api.UserController, accountcontroller *api.AccountController, jwtservice *service.JWTService) *Server {
	return &Server{
		usercontroller:    usercontroller,
		accountcontroller: accountcontroller,
		jwtservice:        jwtservice,
		router:            gin.Default(),
	}
}

func (server *Server) Run(port int) {
	server.router.Run(fmt.Sprintf(":%d", port))
}

func (server *Server) SetupRouter() {
	server.router.POST("/api/v1/Login", server.accountcontroller.Login)
	server.router.POST("/api/v1/SignUp", server.accountcontroller.SignUp)
	server.router.GET("/api/v1/user", server.jwtservice.AuthRequired, server.usercontroller.Users)
	server.router.GET("/api/v1/user/:name", server.jwtservice.AuthRequired, server.usercontroller.User)
	server.router.DELETE("/api/v1/user", server.jwtservice.AuthRequired, server.accountcontroller.Delete)
	server.router.PUT("/api/v1/user", server.jwtservice.AuthRequired, server.accountcontroller.Update)
	docs.SwaggerInfo.BasePath = "/api/v1"
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
