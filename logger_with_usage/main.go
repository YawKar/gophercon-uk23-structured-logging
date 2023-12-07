package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdin, nil))
	slog.SetDefault(logger)

	http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	l := slog.Default().With("path", r.URL.Path) // may appear useful for traceId
	l.Info("start")
	// time=2023-12-07T21:10:02.273+03:00 level=INFO msg=start path=/
	process(w, r, l)
}

func process(w http.ResponseWriter, r *http.Request, l *slog.Logger) {
	l.Info("doing", "item", r.FormValue("item"))
	// time=2023-12-07T21:10:02.273+03:00 level=INFO msg=doing path=/ item=""
}
