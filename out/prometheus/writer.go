package prometheus

import (
	"errors"
	"fmt"
	"github.com/mono83/xray/args"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/mono83/xray"
)

var (
	pType      = []byte("# TYPE ")
	pGauge     = []byte(" gauge\n")
	pCounter   = []byte(" counter\n")
	pHistogram = []byte(" histogram\n")
	pStart     = []byte("{")
	pEnd       = []byte("}")
	pSpace     = []byte(" ")
	pLe        = []byte("le=\"")
	pInf       = []byte("+Inf")
	pQuote     = []byte("\"")
	pEqQuote   = []byte("=\"")
	pComma     = []byte(",")
	pNl        = []byte("\n")

	notAllowedRegex = regexp.MustCompile(`[^a-zA-Z0-9_:]`)
)

// Write outputs contents of exporter into given Writer in prometheus format
func (e *Exporter) Write(w io.Writer) error {
	e.cntWritten++
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

	// Injecting own metrics
	counters = append(
		counters,
		value{value: e.cntHandled, name: "prometheus_exporter_handled"},
		value{value: e.cntWritten, name: "prometheus_exporter_render", args: []xray.Arg{args.Type("writer")}},
		value{value: e.cntHTTP, name: "prometheus_exporter_render", args: []xray.Arg{args.Type("http")}},
	)
	gauges = append(
		gauges,
		value{value: int64(len(e.gauges)), name: "prometheus_exporter_size", args: []xray.Arg{args.Type("gauges")}},
		value{value: int64(len(e.counters)), name: "prometheus_exporter_size", args: []xray.Arg{args.Type("counters")}},
		value{value: int64(len(e.histogram)), name: "prometheus_exporter_size", args: []xray.Arg{args.Type("histogram")}},
	)

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

	// Printing histograms
	for _, v := range histogram {
		if _, err := w.Write(pType); err != nil {
			return err
		}
		if _, err := w.Write(escape(v.name)); err != nil {
			return err
		}
		if _, err := w.Write(pHistogram); err != nil {
			return err
		}

		if err := e.format(w, v.name, "_count", v.args, nil); err != nil {
			return err
		}
		if _, err := w.Write(pSpace); err != nil {
			return err
		}
		if _, err := fmt.Fprint(w, v.counts); err != nil {
			return err
		}
		if _, err := w.Write(pNl); err != nil {
			return err
		}

		if err := e.format(w, v.name, "_sum", v.args, nil); err != nil {
			return err
		}
		if _, err := w.Write(pSpace); err != nil {
			return err
		}
		if _, err := fmt.Fprint(w, v.sum); err != nil {
			return err
		}
		if _, err := w.Write(pNl); err != nil {
			return err
		}

		for i, j := range v.buckets {
			if j == 0 {
				continue
			}
			b := e.buckets[i]
			if err := e.format(w, v.name, "_bucket", v.args, &b); err != nil {
				return err
			}
			if _, err := w.Write(pSpace); err != nil {
				return err
			}
			if _, err := fmt.Fprint(w, j); err != nil {
				return err
			}
			if _, err := w.Write(pNl); err != nil {
				return err
			}
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
	if err := e.format(w, v.name, "", v.args, nil); err != nil {
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

func (e *Exporter) format(w io.Writer, key, suffix string, args []xray.Arg, bucket *int64) error {
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
			if _, err := w.Write(pLe); err != nil {
				return err
			}
			if *bucket == inf {
				if _, err := w.Write(pInf); err != nil {
					return err
				}
			} else {
				if _, err := fmt.Fprint(w, *bucket); err != nil {
					return err
				}
			}
			if _, err := w.Write(pQuote); err != nil {
				return err
			}
			notFirst = true
		}
		for _, a := range e.defaultArgs.Args() {
			if err := writeArg(w, &notFirst, a); err != nil {
				return err
			}
		}
		for _, a := range args {
			if err := writeArg(w, &notFirst, a); err != nil {
				return err
			}
		}
		if _, err := w.Write(pEnd); err != nil {
			return err
		}
	}

	return nil
}

func writeArg(w io.Writer, notFirst *bool, a xray.Arg) error {
	v := a.Value()
	if len(v) == 0 {
		return nil
	}

	if *notFirst {
		if _, err := w.Write(pComma); err != nil {
			return err
		}
	}
	if _, err := w.Write(escape(a.Name())); err != nil {
		return err
	}
	if _, err := w.Write(pEqQuote); err != nil {
		return err
	}
	if _, err := w.Write(escape(a.Value())); err != nil {
		return err
	}
	if _, err := w.Write(pQuote); err != nil {
		return err
	}
	*notFirst = true
	return nil
}

// escape replaces all forbidden characters to underscores
func escape(s string) []byte {
	return []byte(notAllowedRegex.ReplaceAllString(strings.TrimSpace(s), "_"))
}
