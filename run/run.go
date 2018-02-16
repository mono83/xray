package run

import "github.com/mono83/xray"

// Run method runs all provided runnables
func Run(funcs ...Runnable) error {
	ray := xray.BOOT

	for _, r := range funcs {
		if err := r(ray); err != nil {
			return err
		}
	}

	return nil
}
