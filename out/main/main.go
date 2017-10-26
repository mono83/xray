package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/out/os"
	"time"
)

func main() {
	color.NoColor = false

	// Registering stdout printer on ROOT logger
	xray.ROOT.On(os.StdOutLogger(xray.TRACE))

	// Forking
	ray := xray.BOOT.Fork().WithLogger("main-test")

	// Sending all kinds of events
	ray.Trace("This is :name message", args.Name("trace"))
	ray.Debug("This is :name message", args.Name("debug"))
	ray.Info("This is :name message", args.Name("info"))
	ray.Warning("This is :name message", args.Name("warning"))
	ray.Error("This is :name message", args.Name("error"))
	ray.Alert("This is :name message", args.Name("alert"))
	ray.Critical("This is :name message", args.Name("critical"))
	ray.Pass(fmt.Errorf("example passing error"))

	time.Sleep(100 * time.Millisecond)
}
