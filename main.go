package main

import (
	"github.com/hrishin/reservation-system/cmd"
	"log/slog"
	"os"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		slog.Error("failed to execute reservation", "issue", err)
		os.Exit(1)
	}
}
