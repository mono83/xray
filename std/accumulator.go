package std

// AccumulatorInt64 accumulate multiple int64 values
type AccumulatorInt64 struct {
	Depth     int
	Aggregate func([]int64) int64
	values    []int64
	offset    int
	next      *AccumulatorInt64
}

// Add method adds new value to accumulator
func (a *AccumulatorInt64) Add(value int64) {
	if a.Depth < 2 || a.Aggregate == nil {
		return
	}
	if a.values == nil {
		a.values = make([]int64, a.Depth)
	}

	if a.offset == a.Depth {
		// Flushing values to next
		if a.next == nil {
			a.next = &AccumulatorInt64{Depth: a.Depth, Aggregate: a.Aggregate}
		}
		a.next.Add(a.Aggregate(a.values))
		a.offset = 0
	}

	for i := a.Depth - 1; i > 0; i-- {
		a.values[i] = a.values[i-1]
	}
	a.values[0] = value

	a.offset++
}

// List returns all values from accumulator
func (a *AccumulatorInt64) List() []int64 {
	if a.next == nil {
		return a.values[0:a.offset]
	}

	return append(a.values, a.next.List()...)
}
