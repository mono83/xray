package run

import (
	"time"

	"github.com/mono83/xray"
)

// Schedule wraps runnable into another one, which will be invoked at requested rate
func Schedule(r Runnable, interval time.Duration) Runnable {
	return func(ray xray.Ray) error {
		// Initial start
		err := r(ray)

		if err == nil {
			go func() {
				for {
					time.Sleep(interval)
					r(ray.Fork())
				}
			}()
		}

		return err
	}
}
