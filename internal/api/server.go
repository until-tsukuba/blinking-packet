package api

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type server struct {
	engine *gin.Engine
}

// Init initializes web server
func Init() {
	gin.SetMode("debug") // TODO: config

	srv := newServer()
	go srv.run()
}

// newServer generates web server
func newServer() server {
	srv := server{}

	srv.engine = gin.Default()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	srv.engine.Use(cors.New(corsConfig))

	srv.setRoutes()

	return srv
}

// run starts web server
func (srv *server) run() {
	err := srv.engine.Run(":8080")
	if err != nil {
		log.Fatalf("Starting server was failed\n%v", err)
	}
}
