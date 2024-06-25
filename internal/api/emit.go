package api

import (
	"github.com/until-tsukuba/blinking-packet/pkg/emit"
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

	err = emit.EmitPacket(req.PacketType, req.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
