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
	logger := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	slog.SetDefault(logger)
	log.Printf("processed %s of size %d in %s",
		item.Name, item.Size, time.Since(start))
	// will produce smth like:
	// {"time":"2023-12-07T19:06:02.650742433+03:00","level":"INFO","msg":"processed item of size 4 in 600ns"}
	logger.Info("processed item",
		"name", item.Name,
		"size", item.Size,
		"duration", time.Since(start))
	// will produce:
	// {"time":"2023-12-07T19:07:54.054536611+03:00","level":"INFO","msg":"processed item","name":"item","size":4,"duration":127100}
}
