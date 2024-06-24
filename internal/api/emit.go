package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type emitRequest struct {
	PacketType string `json:"packet_type"`
	Value      string `json:"value"`
}

func emitController(ctx *gin.Context) {
	var req emitRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	switch req.PacketType {
	// TODO
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid packet type"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
