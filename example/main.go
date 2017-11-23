package main

import (
	"errors"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/args/env"
	"github.com/mono83/xray/out/os"
)

func main() {
	// Assigning stdout logger to ROOT ray
	// with maximum verbosity
	xray.ROOT.On(os.StdOutLogger(xray.TRACE))

	// Sending message
	xray.ROOT.Info("Hello, world")
	xray.ROOT.Debug("This is debug")
	xray.ROOT.Alert("And this is alert")
	xray.ROOT.Error("Errors are less fancy")
	// Sending message with placeholders
	xray.ROOT.Info("Hello, :name", args.Name("World"))
	// Sending message with multiple placeholders
	// :rayId parameter is always injected automatically. Each ray has own unique ID
	// :pid parameter is provided by `env` packet and contains curren process ID
	xray.ROOT.Info(
		"Hello, :name. I am working on :pid with RayID :rayId. Demo error - :err",
		args.Name("World"),
		args.Error{Err: errors.New("this is error message")},
		env.PID,
	)

	var forked xray.Ray

	// Creating fork - it will have own ID
	forked = xray.ROOT.Fork()
	forked.Info("I have RayID :rayId")

	// Creating fork and setting logger name for it
	// Without this forked rays uses parent's logger name, in current case - "root"
	forked = xray.ROOT.Fork().WithLogger("not root")
	forked.Info("I have RayID :rayId")

	// Assigning placeholder values directly into ray
	// RayID wouldnt change (Fork method not invoked)
	withArgs := forked.With(args.Name("foo"), args.ID64(42))
	withArgs.Info("I have RayID :rayId, name :name and ID :id")

	// RayID can be changed manually
	// This is useful on SOA, when one service communicates with other - client can
	// send it's Ray ID to server and server can it set manually to preserve
	// inter-service constistent logging
	withCustom := forked.WithRayID("http-request-123456")
	withCustom.Info("I have RayID :rayId")

	// Waiting until all output is done
	time.Sleep(time.Millisecond * 50)
}
