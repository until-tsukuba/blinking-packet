package main

import (
	"github.com/until-tsukuba/blinking-packet/internal/api"
	"github.com/until-tsukuba/blinking-packet/internal/dump"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	api.Init()
	go dump.DumpAndBlink("eth0", 60, 20)

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	_ = <-trap
}
