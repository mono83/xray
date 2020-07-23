package std

import (
	"strings"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
)

// queryType function attempts to recognize SQL query type
func queryType(sql string) string {
	sql = strings.TrimSpace(sql)
	if len(sql) < 6 {
		return ""
	}

	switch sql[0:6] {
	case "select":
		return "select"
	case "insert":
		return "insert"
	case "update":
		return "update"
	case "delete":
		return "delete"
	default:
		return ""
	}
}

// SQLLog command performs SQL logging
func SQLLog(r xray.Ray, sql string, elapsed xray.NanoHolder, err error) {
	if r == nil {
		return
	}

	// Sending SQL latency
	r.Duration("sql.latency", elapsed, args.SQL(sql))

	list := []xray.Arg{args.SQL(sql), args.Delta(time.Duration(elapsed.Nanoseconds()))}
	if t := queryType(sql); len(t) > 0 {
		list = append(list, args.Type(t))
	}

	// Sending message
	if err == nil {
		r.Debug("Query :sql done in :delta", list...)
	} else {
		list = append(list, args.Error{Err: err})
		r.Error("Query :sql done in :delta with error :err", list...)
	}
}

// WrapSQLQuery method takes SQL reader function and returns wrapped one, that will
// supply additional logging data into ray
func WrapSQLQuery(r xray.Ray, f func(interface{}, string, ...interface{}) error) func(interface{}, string, ...interface{}) error {
	if r == nil {
		return f
	}

	return func(target interface{}, sql string, sqlargs ...interface{}) error {
		// Starting timer
		now := time.Now()

		// Performing request
		err := f(target, sql, sqlargs...)

		// Logging
		SQLLog(r, sql, time.Now().Sub(now), err)

		// Response
		return err
	}
}
