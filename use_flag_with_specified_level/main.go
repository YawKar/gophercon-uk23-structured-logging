package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"
)

type Item struct {
	Name string
	Size int
}

var level slog.Level

func main() {
	start := time.Now()
	item := Item{"item", 4}

	// go run . -level {INFO | WARN | DEBUG | ERROR}
	flag.TextVar(&level, "level", level, "log level")
	flag.Parse()

	opts := &slog.HandlerOptions{Level: level}

	logger := slog.New(slog.NewTextHandler(os.Stdin, opts))
	slog.SetDefault(logger)
	slog.LogAttrs(context.TODO(), slog.LevelInfo.Level(), "processed item",
		slog.String("name", item.Name),
		slog.Int("size", item.Size),
		slog.Duration("duration", time.Since(start)))
	// will produce if without level option (INFO level is default)
	// and will not if, for instance, use WARN:
	// time=2023-12-07T19:19:05.696+03:00 level=INFO msg="processed item" name=item size=4 duration=12µs
}
