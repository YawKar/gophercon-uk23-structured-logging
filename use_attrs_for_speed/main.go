package main

import (
	"context"
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
	slog.LogAttrs(context.TODO(), slog.LevelInfo.Level(), "processed item",
		slog.String("name", item.Name),
		slog.Int("size", item.Size),
		slog.Duration("duration", time.Since(start)))
	// will produce:
	// time=2023-12-07T19:19:05.696+03:00 level=INFO msg="processed item" name=item size=4 duration=12µs
}
