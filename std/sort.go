package std

// Int64Sorter is wrapper for slice of int64 values, that implements sort.Interface
type Int64Sorter []int64

func (i Int64Sorter) Len() int           { return len(i) }
func (i Int64Sorter) Swap(x, y int)      { i[x], i[y] = i[y], i[x] }
func (i Int64Sorter) Less(x, y int) bool { return i[x] < i[y] }
