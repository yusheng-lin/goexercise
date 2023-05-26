package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	server, cleanup, err := initapp()
	defer cleanup()

	if err != nil {
		log.Fatal().Err(err)
		return
	}
	server.SetupRouter()
	server.Run(8080)
}
