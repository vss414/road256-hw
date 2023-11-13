package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func New(filename string) (*logrus.Logger, error) {
	logger := logrus.New()

	logFilePath := filepath.Join("logs", filename)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", logFilePath, err)
	}

	if err := os.Chmod(logFilePath, 0600); err != nil {
		return nil, fmt.Errorf("failed to chmod log file %q: %w", logFilePath, err)
	}

	mw := io.MultiWriter(f)

	logger.SetReportCaller(true)
	logger.SetOutput(mw)
	logger.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	return logger, err
}
