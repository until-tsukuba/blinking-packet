package main

import (
	"github.com/until-tsukuba/blinking-packet/internal/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	api.Init()

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	_ = <-trap
}
