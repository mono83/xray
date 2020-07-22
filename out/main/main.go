package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/out/os"
)

func main() {
	color.NoColor = false

	// Registering stdout printer on ROOT logger
	xray.ROOT.On(os.StdOutLogger(xray.TRACE))
	xray.ROOT.On(os.StdOutDumper())

	// Forking
	ray := xray.BOOT.Fork().WithLogger("main-test")

	// Sending all kinds of events
	ray.Trace("This is :name message", args.Name("trace"))
	ray.Debug("This is :name message", args.Name("debug"))
	ray.Info("This is :name message in ray :rayId", args.Name("info"))
	ray.Warning("This is :name message", args.Name("warning"))
	ray.Error("This is :name message", args.Name("error"))
	ray.Alert("This is :name message", args.Name("alert"))
	ray.Critical("This is :name message", args.Name("critical"))
	ray.Pass(fmt.Errorf("example passing error"))

	// Sending dump information
	ray.OutBytes([]byte("Hello, world"))
	ray.InBytes([]byte("Received response"))

	// Ray ID check
	ray.Info("Send in own with :rayId")
	xray.BOOT.Info("Send in BOOT with :rayId")
	xray.ROOT.Info("Send in ROOT with :rayId")

	time.Sleep(100 * time.Millisecond)
}
