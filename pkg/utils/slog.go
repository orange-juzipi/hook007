package utils

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func InitSlog() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// update time
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{
					Key:   "time",
					Value: slog.AnyValue(time.Now().Format(time.DateTime)),
				}
			}

			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}

			return a
		},
	}))

	slog.SetDefault(logger)
}
