package main

import (
	"log"
	"log/slog"
	"os"
	"time"
)

type Item struct {
	Name string
	Size int
}

func main() {
	start := time.Now()
	item := Item{"item", 4}
	logger := slog.New(slog.NewTextHandler(os.Stdin, nil))
	slog.SetDefault(logger)
	log.Printf("processed %s of size %d in %s",
		item.Name, item.Size, time.Since(start))
	// will produce smth like:
	// time=2023-12-07T19:02:37.026+03:00 level=INFO msg="processed item of size 4 in 500ns"
}
