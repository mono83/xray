package xray

import (
	"strings"

	"github.com/mono83/xray/args"
)

func (r *ray) Pass(err error) error {
	if err != nil {
		r.Error("Error encountered - :err", args.Error{Err: err})
	}

	return err
}

func (r *ray) PassS(message string, err error) error {
	if err != nil {
		if strings.Contains(message, ":err") {
			r.Error(message, args.Error{Err: err})
		} else {
			r.Error(message+" :err", args.Error{Err: err})
		}
	}

	return err
}

func (r *ray) OutBytes(data []byte, args ...Arg) { r.Emit(newDumpingEvent(r, OUT, data, args)) }
func (r *ray) InBytes(data []byte, args ...Arg)  { r.Emit(newDumpingEvent(r, IN, data, args)) }

func (r *ray) Trace(msg string, args ...Arg)    { r.Emit(newLogEvent(r, TRACE, msg, args)) }
func (r *ray) Debug(msg string, args ...Arg)    { r.Emit(newLogEvent(r, DEBUG, msg, args)) }
func (r *ray) Info(msg string, args ...Arg)     { r.Emit(newLogEvent(r, INFO, msg, args)) }
func (r *ray) Warning(msg string, args ...Arg)  { r.Emit(newLogEvent(r, WARNING, msg, args)) }
func (r *ray) Error(msg string, args ...Arg)    { r.Emit(newLogEvent(r, ERROR, msg, args)) }
func (r *ray) Alert(msg string, args ...Arg)    { r.Emit(newLogEvent(r, ALERT, msg, args)) }
func (r *ray) Critical(msg string, args ...Arg) { r.Emit(newLogEvent(r, CRITICAL, msg, args)) }

func (r *ray) Inc(key string, args ...Arg) { r.Emit(newMetricEvent(r, INCREMENT, key, 1, args)) }
func (r *ray) Increment(key string, value int64, args ...Arg) {
	r.Emit(newMetricEvent(r, INCREMENT, key, value, args))
}
func (r *ray) Gauge(key string, value int64, args ...Arg) {
	r.Emit(newMetricEvent(r, GAUGE, key, value, args))
}
func (r *ray) Duration(key string, value NanoHolder, args ...Arg) {
	r.Emit(newMetricEvent(r, DURATION, key, value.Nanoseconds(), args))
}
