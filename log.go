package main

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	writer := io.MultiWriter(
		&lumberjack.Logger{
			Filename:   "app.log",
			MaxAge:     30,  // days
			MaxSize:    256, // megabytes
			MaxBackups: 128, // files
		},
		io.MultiWriter(os.Stdout),
	)
	opt := &slog.HandlerOptions{
		AddSource: true,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(writer, opt)))
}
