package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if _, err := maxprocs.Set(maxprocs.Logger(logger.Info)); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			logger.Info("timeout")
			return
		case <-ticker.C:
			logger.Info("tick")
		}
	}
}
