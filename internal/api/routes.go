package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// setRoutes sets routing
func (srv *server) setRoutes() {
	engine := srv.engine

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	engine.POST("/emit", emitController)
}
