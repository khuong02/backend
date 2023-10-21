package logger

import (
	"io"
	"log"
	"os"

	"golang.org/x/exp/slog"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(path string, level int) (*Logger, error) {
	writer, err := CreateHybridWriter(path)
	if err != nil {
		return nil, err
	}
	handler := slog.NewJSONHandler(writer,
		&slog.HandlerOptions{
			AddSource: false,
			Level:     slog.Level(level),
		})
	logger := slog.New(&CustomHandler{
		Handler: handler,
	})
	return &Logger{
		logger,
	}, nil
}

func CreateHybridWriter(path string) (io.Writer, error) {
	fileWriter, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		return nil, err
	}
	log.Default()
	cliWrite := os.Stderr
	return io.MultiWriter(fileWriter, cliWrite), nil
}
