package xray

// Wrap function wraps provided ray, injecting logger and arguments
// If nil provided, it creates new ray by forking xray.ROOT
func Wrap(ray Ray, logger string, args ...Arg) Ray {
	if ray == nil {
		ray = ROOT.Fork()
	}

	return ray.WithLogger(logger).With(args...)
}
