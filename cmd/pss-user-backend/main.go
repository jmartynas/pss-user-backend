package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logger, err := newLogger()
	if err != nil {
		fmt.Println("Could not create logger", err)
		os.Exit(1)
	}
	_ = logger

	// create mysql connection

	// create redis connection

	// create handlers

	// create server
	srv := http.Server{
		Addr: os.Getenv("PSS_ADDRESS"),
	}

	// create server closer

	// run server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("server error", slog.String("error", err.Error()))
	}
}

func newLogger() (*slog.Logger, error) {
	b, _ := strconv.ParseBool(os.Getenv("PSS_LOG_TO_FILE"))

	var logger *slog.Logger
	if b {
		logFile, err := os.OpenFile(
			"/var/log/pss-user-backend.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0755,
		)
		if err != nil {
			return nil, err
		}

		logger = slog.New(slog.NewJSONHandler(logFile, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return logger, nil
}
