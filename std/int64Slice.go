package std

import "sort"

// Int64Slice is wrapper for slice of int64 values, that implements sort.Interface
type Int64Slice []int64

func (i Int64Slice) Len() int           { return len(i) }
func (i Int64Slice) Swap(x, y int)      { i[x], i[y] = i[y], i[x] }
func (i Int64Slice) Less(x, y int) bool { return i[x] < i[y] }

// Sort sorts underlying slice in ascending order
func (i Int64Slice) Sort() { sort.Sort(i) }

// Analyze return min, max, avg, sum and count values from underlying slice
func (i Int64Slice) Analyze() (count int, min, max, avg, sum int64) {
	count = len(i)
	if count == 0 {
		return
	} else if count == 1 {
		min = i[0]
		max = min
		avg = min
		sum = min
		return
	}

	for k, v := range i {
		if k == 0 {
			min = v
			max = v
		} else {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}

		sum += v
	}

	avg = sum / int64(count)
	return
}

// Percentile returns percentile value itself and slice of values before it (inclusive)
func (i Int64Slice) Percentile(p int) (value int64, slice []int64) {
	if p < 1 {
		p = 1
	} else if p > 99 {
		p = 99
	}
	index := len(i) * p / 100
	value = i[index]
	if index == len(i)-1 {
		slice = i
	} else {
		slice = i[0 : index+1]
	}
	return
}
