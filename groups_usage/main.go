package main

import (
	"log/slog"
	"os"
)

type Item struct {
	Name string
	Size int
}

func renHoek(l *slog.Logger) {
	l.Info("msg", slog.Group("name",
		"first", "Ren",
		"last", "Hoek"))
	// {"time":"2023-12-07T20:54:31.251031355+03:00","level":"INFO","msg":"msg","name":{"first":"Ren","last":"Hoek"}}
}

func thing1(l *slog.Logger) {
	l.Info("thing1 logged", "value", 5)
	// {"time":"2023-12-07T20:58:03.108319801+03:00","level":"INFO","msg":"thing1 logged","thing1":{"value":5}}
}

func thing2(l *slog.Logger) {
	l.Info("thing2 logged", "value", "high")
	// {"time":"2023-12-07T20:58:03.108339601+03:00","level":"INFO","msg":"thing2 logged","thing2":{"value":"high"}}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	renHoek(logger)
	thing1(logger.WithGroup("thing1"))
	thing2(logger.WithGroup("thing2"))
}
