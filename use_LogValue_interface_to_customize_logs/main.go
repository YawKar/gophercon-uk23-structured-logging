package main

import (
	"log/slog"
	"os"
)

type Name struct {
	First, Last string
}

func (n Name) LogValue() slog.Value { // slog.Value|s are useful in case we use different handlers (YAML, TOML and etc)
	return slog.GroupValue(
		slog.String("first", n.First),
		slog.String("last", n.Last),
	)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	logger.Info("the message", "the name", Name{"first", "last"})
	// {"time":"2023-12-07T21:20:13.246124336+03:00","level":"INFO","msg":"the message","the name":{"first":"first","last":"last"}}
}
