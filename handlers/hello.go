package handlers

import (
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	defer txn.StartSegment("hello-segment").End()

	w.Write([]byte("Hello, World!"))
}
