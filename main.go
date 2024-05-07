package main

import (
	"fmt"
	"github.com/hrishin/reservation-system/cmd"
	"github.com/hrishin/reservation-system/internal/state"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reservation",
	Short: "A CLI application for booking and cancellation of flights",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func iniCmd(stare state.Storable) {
	rootCmd.AddCommand(cmd.NewBookingCommand(stare))
	rootCmd.AddCommand(cmd.NewCancelCommand(stare))
}

func main() {
	iniCmd(state.NewFileSate(""))
	Execute()
}
