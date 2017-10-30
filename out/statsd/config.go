package statsd

import (
	"errors"
	"github.com/mono83/udpwriter"
	"github.com/mono83/xray"
	"github.com/mono83/xray/out"
	"net"
	"time"
)

// Config holds configuration for StatsD client
type Config struct {
	Address       string   `json:"address" yaml:"address" toml:"address"`
	Buffer        int      `json:"bufferMillis" yaml:"bufferMillis" toml:"bufferMillis"`
	Args          bool     `json:"args" yaml:"args" toml:"args"`
	ArgsWhiteList []string `json:"argsWhiteList" yaml:"argsWhiteList" toml:"argsWhiteList"`
	ArgsBlackList []string `json:"argsBlackList" yaml:"argsBlackList" toml:"argsBlackList"`
}

// Validate validates configuration contents
func (c Config) Validate() error {
	if len(c.Address) == 0 {
		return errors.New("empty StatsD binding address")
	}
	if c.Buffer < 100 {
		return errors.New("at least 100ms buffering must be configured")
	}

	return nil
}

// Raw method builds synchronous sender for StatsD.
// It is not recommended to use this method, call Build() instead
func (c Config) Raw() (xray.Handler, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	netAddr, err := net.ResolveUDPAddr("udp", c.Address)
	if err != nil {
		return nil, err
	}

	// Building sender
	send := &sender{
		target:     udpwriter.New(netAddr),
		argAllowed: c.Args,
		argFilter:  xray.ArgFilterDoubleList(c.ArgsWhiteList, c.ArgsBlackList),
	}

	return send.handle, nil
}

// MustBuild is an alias for Build but panics on error
func (c Config) MustBuild() xray.Handler {
	hld, err := c.Build()
	if err != nil {
		panic(err)
	}

	return hld
}

// Build builds asynchronous buffered StatsD receiver
func (c Config) Build() (xray.Handler, error) {
	hld, err := c.Raw()
	if err != nil {
		return nil, err
	}

	return out.Filter(
		out.Buffer(
			out.Splitter(
				hld,
				10,
			),
			time.Duration(c.Buffer)*time.Millisecond,
		),
		func(event xray.Event) bool {
			if event == nil {
				return false
			}
			_, ok := event.(xray.MetricsEvent)
			return ok
		},
	), nil
}
