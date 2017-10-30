package xray

// ArgFilter is function, that can be used in out sender to limit arguments,
// that should be passed to output
type ArgFilter func([]Arg) []Arg

func argFilterAllowAll(in []Arg) []Arg {
	return in
}

func argFilterDenyAll(in []Arg) []Arg {
	return nil
}

// ArgFilterBlacklist returns function, that can be used to blacklist certain args by name
func ArgFilterBlacklist(names []string) ArgFilter {
	if len(names) == 0 {
		return argFilterAllowAll
	}

	m := map[string]bool{}
	for _, v := range names {
		m[v] = true
	}

	return func(args []Arg) []Arg {
		if len(args) == 0 {
			return nil
		}
		response := []Arg{}
		for _, v := range args {
			if _, ok := m[v.Name()]; !ok {
				response = append(response, v)
			}
		}
		return response
	}
}

// ArgFilterWhitelist returns function, that performs args whitelisting by name
func ArgFilterWhitelist(names []string) ArgFilter {
	if len(names) == 0 {
		return argFilterDenyAll
	}

	m := map[string]bool{}
	for _, v := range names {
		m[v] = true
	}

	return func(args []Arg) []Arg {
		if len(args) == 0 {
			return nil
		}
		response := []Arg{}
		for _, v := range args {
			if _, ok := m[v.Name()]; ok {
				response = append(response, v)
			}
		}
		return response
	}
}

// ArgFilterDoubleList determines white or blacklist argument filtering.
// Whitelisting is preferred and has higher priority
func ArgFilterDoubleList(whitelist, blacklist []string) ArgFilter {
	if len(whitelist) > 0 {
		return ArgFilterWhitelist(whitelist)
	}

	return ArgFilterBlacklist(blacklist)
}
