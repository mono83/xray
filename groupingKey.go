package xray

import (
	"bytes"
	"sort"
)

// MetricGroupingKey calculates grouping key for metrics event.
// Generated value can be used as map keys in buffers and aggregations
func MetricGroupingKey(e MetricsEvent, f ArgFilter) string {
	if e == nil {
		return ""
	}

	args := e.Args()
	if f != nil {
		args = f(args)
	}

	return e.GetKey() + "\t" + GroupingKeyArgs(args)
}

// GroupingKeyArgs builds grouping key (some kind of fingerprint) for provided arguments slice
func GroupingKeyArgs(args []Arg) string {
	l := len(args)
	if l == 0 {
		return ""
	} else if l == 1 {
		return args[0].Name() + "\t" + args[0].Value()
	}

	list := make([]stuple, l)
	for i, v := range args {
		list[i] = stuple{name: v.Name(), value: v.Value()}
	}

	sort.Sort(stuplelist(list))
	buf := bytes.NewBuffer(nil)
	for i, v := range list {
		if i > 0 {
			buf.WriteRune('\t')
		}
		buf.WriteString(v.name)
		buf.WriteRune('\t')
		buf.WriteString(v.value)
	}

	return buf.String()
}

// stuple is a string tuple
type stuple struct {
	name, value string
}

// stuplelist is slice of string tuple values
type stuplelist []stuple

// Len is the number of elements in the collection.
func (s stuplelist) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s stuplelist) Less(i, j int) bool {
	return s[i].name < s[j].name
}

// Swap swaps the elements with indexes i and j.
func (s stuplelist) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
