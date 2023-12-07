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

var level = &slog.LevelVar{}

func main() {
	start := time.Now()
	item := Item{"item", 4}

	// go run . -level {INFO | WARN | DEBUG | ERROR}
	flag.TextVar(level, "level", level, "log level")
	flag.Parse()

	opts := &slog.HandlerOptions{Level: level}

	logger := slog.New(slog.NewTextHandler(os.Stdin, opts))
	slog.SetDefault(logger)

	level.Set(slog.LevelDebug) // Dynamically setting log level to debug, so log will be produced

	slog.LogAttrs(context.TODO(), slog.LevelDebug.Level(), "processed item",
		slog.String("name", item.Name),
		slog.Int("size", item.Size),
		slog.Duration("duration", time.Since(start)))
}
