X-Ray
=====

X-Ray is logging framework, created for concurrent environment. Central idea - loggers forms 
tree structure - starting from `xray.ROOT` and then forking to arbitrary level with full
properties inheritance.

Loggers inside framework are named **rays**. Each forked ray will inherit all properties from 
parent, including assigned output adapters. 

Examples are located in `github.com/mono83/xray/example` package

## Ray object

### Ray manipulation methods

| Name                       | Description                                                                                                                                            |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| `GetRayID()`               | Return `string` ID of current Ray                                                                                                                      |
| `GetLogger()`              | Returns name of logger, `string`                                                                                                                       |
| `GetMetricPrefix()`        | Returns `string` prefix, that will be attached to all metric events                                                                                    |
| `GetArguments()`           | Returns all arguments (placeholder values), attached to ray as `xray.Bucket`                                                                           |
| `Fork()`                   | Builds and returns new `xray.Ray`. New object will inherit all propertes of original, but with new rayID                                               |
| `WithRayID(string)`        | Clones ray and sets provided ID to it                                                                                                                  |
| `WithLogger(string)`       | Clones ray and sets provided logger name to it                                                                                                         |
| `WithMetricPrefix(string)` | Clones ray and **appends** metric prefix to it                                                                                                         |
| `With(...xray.Arg)`        | Clones ray and **appends** provided arguments to clone                                                                                                 |
| `On(xray.Handler)`         | Attaches `Handler`, that will receive all logging and metrics event. Registered handler will received events from ray itself and all its forked childs |
| `Emit(...xray.Event)`      | Emits arbitrary `xray.Event` to ray                                                                                                                    |

### Logging methods

| Method                                   | Description                                                                                                                                              |
|------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Trace(msg string, args ...xray.Arg)`    | Emits log event with `xray.TRACE` level                                                                                                                  |
| `Debug(msg string, args ...xray.Arg)`    | Emits log event with `xray.DEBUG` level                                                                                                                  |
| `Info(msg string, args ...xray.Arg)`     | Emits log event with `xray.INFO` level                                                                                                                   |
| `Warning(msg string, args ...xray.Arg)`  | Emits log event with `xray.WARNING` level                                                                                                                |
| `Error(msg string, args ...xray.Arg)`    | Emits log event with `xray.ERROR` level                                                                                                                  |
| `Alert(msg string, args ...xray.Arg)`    | Emits log event with `xray.ALERT` level                                                                                                                  |
| `Critical(msg string, args ...xray.Arg)` | Emits log event with `xray.CRITICAL` level                                                                                                               |
| `Pass(error) error`                      | If provided `error` not nil, emits logging event `xray.ERROR` level. In any case this methods returns error, provided to it.                             |
| `PassS(msg string, err error) error`     | If provided `error` not nil, emits logging event `xray.ERROR` level with provided string prefix. In any case this methods returns error, provided to it. |

### Metric reporting methods

| Method                                      | Description                                                                                                                                                              |
|---------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Increment(string, int64, ...xray.Arg)`     | Emits metric event with `xray.INCREMENT` metric type and `int64` incrementation value                                                                                    |
| `Inc(string, ...xray.Arg)`                  | Emits metric event with `xray.INCREMENT` metric type and `1` as incrementation value                                                                                     |
| `Gauge(string, int64, ...xray.Arg)`         | Emits metric event with `xray.GAUGE` metric type and `int64` gauge value                                                                                                 |
| `Duration(string, NanoHolder, ...xray.Arg)` | Emits metric event with `xray.DURATION` metric type. `NanoHolder` is interface with `Nanoseconds() int64` method, that is implemented even in `time.Duration` structure. |


## Arguments

In general case, arguments a objects, used for message interpolation. In special cases, like **Logstash** or **InfluxDB** this arguments can be sent to remote servers.

Argument is pretty simple and described by interface:

```go
// Arg describes ray logging qualifier (argument)
type Arg interface {
	// Name returns argument key (name)
	Name() string
	// Value returns string representation of argument value
	Value() string
	// Scalar returns raw representation of argument value. It can be scalar value or slice of scalar values. 
	Scalar() interface{}
}
```

Predefined arguments are located in `github.com/mono83/xray/args` package and it subpackages:

| Argument       | Name      | Scalar   | Instantiation                                                         |
|----------------|-----------|----------|-----------------------------------------------------------------------|
| `args.Nil`     | *any*     | `nil`    | Cast from `string`, that will be used as Name                         |
| `args.Error`   | `"err"`   | `string` | `args.Error{Err: <error value>}`                                      |
| `args.Int`     | *any*     | `int`    | `args.Int{N: <string name>, V: <int value>}`                          |
| `args.Int64`   | *any*     | `int64`  | `args.Int64{N: <string name>, V: <int64 value>}`                      |
| `args.Count`   | `"count"` | `int`    | Cast from `int`, that will be used as Value                           |
| `args.ID64`    | `"id"`    | `int64`  | Cast from `int64`, that will be used as Value                         |
| `args.String`  | *any*     | `string` | `args.String{N: <string name>, V: <str value>}`                       |
| `args.Name`    | `"name"`  | `string` | Cast from `string`, that will be used as Value                        |
| `args.Type`    | `"type"`  | `string` | Cast from `string`, that will be used as Value                        |
| `args.AppName` | `"app"`   | `string` | Cast from `string`, that will be used as Value                        |
| `args.URL`     | `"url"`   | `string` | Cast from `string`, that will be used as Value                        |
| `args.SQL`     | `"sql"`   | `string` | Cast from `string`, that will be used as Value                        |
| `args.Delta`   | `"delta"` | `int64`  | Cast from `time.Duration`. Method `Scalar()` will returns nanoseconds |

There are also special arguments, that wrap slices 

| Argument        | Name   | Values List |
|-----------------|--------|-------------|
| `args.ID64List` | `"id"` | `[]int64`   |
| `args.NameList` | "name" | `[]string`  |

And even more special arguments, that are instantiated by default and located in `github.com/mono83/args/env`

| Argument            | Instance         | Name         | Scalar   | Description                                                |
|---------------------|------------------|--------------|----------|------------------------------------------------------------|
| `env.ArgPID` 	      | `env.PID`        | `"pid"`      | `int`    | Contains process ID (pid) of current application           |
| `env.ArgHostName`   | `env.HostName`   | `"hostname"` | `string` | Contains host name of machine, this application running on |
| `env.ArgSystemUser` | `env.SystemUser` | `"username"` | `string` | Contain information of user, invoked application           |

## Prometheus exporter

Xray is bundled with data exporter, compatible with Prometheus. To enable, attach it to `xray.ROOT` or other ray:

```go
import "github.com/mono83/xray/out/prometheus"

// Declaring exporter
prom := prometheus.NewExporter(
        nil, // Metrics arguments filtering func
        nil, // Default arguments to inject
        time.Millisecond, // Time buckets for latency histograms...
        10 * time.Millisecond,
        50 * time.Millisecond,
        100 * time.Millisecond,
        250 * time.Millisecons,
        500 * time.Millisecond, 
        time.Second,
        10 * time.Second,
)

// Registering on ROOT ray
xray.ROOT.On(prom.Handle)

// Starting HTTP server, that will return metrics on any URL
panic(http.ListenAndServe(":5557", prom))
```

In addition to served metrics, this exporter also sends its own to diagnose performance issues:

* `prometheus_exporter_handled` - amount of handled metrics events, counter.
* `prometheus_exporter_render` - amount of render requests, counter. `type="writer"` stands for any writer invocation,
  `type="http"` shows only HTTP rendering requests.
* `prometheus_exporter_size` - amount of items in exporter, gauge. `type="gauges"` calculates only gauges,
  `type="counters"` is for counters and `type="histogram"` is for histograms.  