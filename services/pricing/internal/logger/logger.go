package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

// Init initializes the global logger
func Init() {
	Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}
