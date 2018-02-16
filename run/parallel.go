package run

import (
	"github.com/mono83/xray"
	"sync"
)

// Parallel method wraps multiple multiple runnables into one
// Wrapped entries will be invoked in parallel
func Parallel(list ...Runnable) Runnable {
	return func(ray xray.Ray) (err error) {
		wg := sync.WaitGroup{}
		wg.Add(len(list))
		for _, r := range list {
			go func(x Runnable) {
				if e := x(ray); e != nil {
					err = e
				}
				wg.Done()
			}(r)
		}
		wg.Wait()
		return
	}
}
