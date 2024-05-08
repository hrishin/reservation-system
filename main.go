package main

import (
	"fmt"
	"github.com/hrishin/reservation-system/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %v\n", err)
		os.Exit(1)
	}
}
