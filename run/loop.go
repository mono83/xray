package run

import (
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Loop is infinite loop
func Loop(xray.Ray) error {
	for {
		time.Sleep(time.Hour)
	}
}

// SigBreakLoop starts infinite loop, that can be broken on SIGTERM or SIGINT
func SigBreakLoop(ray xray.Ray) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ray.Info("Registered signal handlers for SIGINT and SIGTERM")

	s := <-c
	ray.Info("Got signal :name", args.Name(s.String()))
	return nil
}
