package xray

import "time"

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
func (r *ray) Duration(key string, value time.Duration, args ...Arg) {
	r.Emit(newMetricEvent(r, DURATION, key, value.Nanoseconds(), args))
}
