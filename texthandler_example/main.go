package main

import (
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
	logger.Info("processed item",
		"name", item.Name,
		"size", item.Size,
		"duration", time.Since(start))
	// will produce smth like:
	// time=2023-12-07T18:57:25.817+03:00 level=INFO msg="processed item" name=item size=4 duration=400ns
}
