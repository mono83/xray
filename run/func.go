package run

import "github.com/mono83/xray"

// Func wraps common Go function into Runnable
func Func(f func()) Runnable {
	return func(ray xray.Ray) error {
		f()
		return nil
	}
}

// FuncE wraps common Go function, able to return error, into Runnable
func FuncE(f func() error) Runnable {
	return func(ray xray.Ray) error {
		return f()
	}
}
