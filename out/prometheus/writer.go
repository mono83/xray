package prometheus

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/mono83/xray"
)

var (
	pType    = []byte("# TYPE ")
	pGauge   = []byte(" gauge\n")
	pCounter = []byte(" counter\n")
	pStart   = []byte("{")
	pEnd     = []byte("}")
	pSpace   = []byte(" ")
	pLe      = []byte("le=\"")
	pInf     = []byte("+Inf")
	pQuote   = []byte("\"")
	pEqQuote = []byte("=\"")
	pComma   = []byte(",")
	pNl      = []byte("\n")

	notAllowedRegex = regexp.MustCompile(`[^a-zA-Z0-9_:]`)
)

// Write outputs contents of exporter into given Writer in prometheus format
func (e *Exporter) Write(w io.Writer) error {
	if w == nil {
		return errors.New("empty writer")
	}

	// Copying data into local slices to avoid
	// IO impact by given writer
	e.mutex.Lock()
	var i int
	counters := make([]value, len(e.counters))
	gauges := make([]value, len(e.gauges))
	histogram := make([]timeValue, len(e.histogram))

	i = 0
	for _, v := range e.counters {
		counters[i] = *v
		i++
	}
	i = 0
	for _, v := range e.gauges {
		gauges[i] = *v
		i++
	}
	i = 0
	for _, v := range e.histogram {
		histogram[i] = *v
		i++
	}
	e.mutex.Unlock()

	// Sorting in alphabetical order
	sort.Slice(counters, func(i, j int) bool {
		return counters[i].name < counters[j].name
	})
	sort.Slice(gauges, func(i, j int) bool {
		return gauges[i].name < gauges[j].name
	})
	sort.Slice(histogram, func(i, j int) bool {
		return histogram[i].name < histogram[j].name
	})

	// Printing gauges
	for _, v := range gauges {
		if err := e.writeCounterOrGauge(w, pGauge, v); err != nil {
			return err
		}
	}

	// Printing counters
	for _, v := range counters {
		if err := e.writeCounterOrGauge(w, pCounter, v); err != nil {
			return err
		}
	}

	return nil
}

func (e *Exporter) writeCounterOrGauge(w io.Writer, t []byte, v value) error {
	if _, err := w.Write(pType); err != nil {
		return err
	}
	if _, err := w.Write(escape(v.name)); err != nil {
		return err
	}
	if _, err := w.Write(t); err != nil {
		return err
	}
	if err := e.format(v.name, "", v.args, nil, w); err != nil {
		return err
	}
	if _, err := w.Write(pSpace); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, v.value); err != nil {
		return err
	}
	if _, err := w.Write(pNl); err != nil {
		return err
	}

	return nil
}

func (e *Exporter) format(key, suffix string, args []xray.Arg, bucket *int64, w io.Writer) error {
	if _, err := w.Write(escape(key)); err != nil {
		return err
	}
	if len(suffix) > 0 {
		if _, err := w.Write([]byte(suffix)); err != nil {
			return err
		}
	}
	if len(args) > 0 || e.defaultArgs.Size() > 0 || bucket != nil {
		if _, err := w.Write(pStart); err != nil {
			return err
		}
		notFirst := false
		if bucket != nil {
			w.Write(pLe)
			if *bucket == inf {
				w.Write(pInf)
			} else {
				fmt.Fprint(w, pInf)
			}
			w.Write(pQuote)
			notFirst = true
		}
		for _, a := range e.defaultArgs.Args() {
			v := a.Value()
			if len(v) == 0 {
				continue
			}

			if notFirst {
				w.Write(pComma)
			}
			w.Write(escape(a.Name()))
			w.Write(pEqQuote)
			w.Write(escape(a.Value()))
			w.Write(pQuote)
			notFirst = true
		}
		for _, a := range args {
			v := a.Value()
			if len(v) == 0 {
				continue
			}

			if notFirst {
				w.Write(pComma)
			}
			w.Write(escape(a.Name()))
			w.Write(pEqQuote)
			w.Write(escape(a.Value()))
			w.Write(pQuote)
			notFirst = true
		}
		w.Write(pEnd)
	}

	return nil
}

// escape replaces all forbidden characters to underscores
func escape(s string) []byte {
	return []byte(notAllowedRegex.ReplaceAllString(strings.TrimSpace(s), "_"))
}
