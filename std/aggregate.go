package std

// AggregateSum is aggregation function, that returns sum of provided values
func AggregateSum(int64s []int64) int64 {
	sum := int64(0)
	for _, v := range int64s {
		sum += v
	}
	return sum
}

// AggregateAvg is aggregation function, that returns average value
func AggregateAvg(int64s []int64) int64 {
	if len(int64s) == 0 {
		return 0
	}

	return AggregateSum(int64s) / int64(len(int64s))
}

// AggregateMax is aggregation function, that returns max value from list
func AggregateMax(int64s []int64) int64 {
	var max int64
	for k, v := range int64s {
		if k == 0 || v > max {
			max = v
		}
	}
	return max
}

// AggregateMin is aggregation function, that returns min value from list
func AggregateMin(int64s []int64) int64 {
	var min int64
	for k, v := range int64s {
		if k == 0 || v < min {
			min = v
		}
	}
	return min
}
