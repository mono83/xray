package prometheus

import "net/http"

// ServeHTTP is http.Handler interface implementation
func (e *Exporter) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	e.cntHTTP++
	w.Header().Set("Content-Type", "text/plain")
	_ = e.Write(w)
}
