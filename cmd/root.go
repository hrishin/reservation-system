package cmd

import (
	"github.com/hrishin/reservation-system/internal/state"
	"github.com/spf13/cobra"
)

var statePath string

var rootBookingState state.Storable

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "booking",
		Short: "A CLI application for booking and cancellation of flight reservations",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVar(&statePath, "state-file", "", "directory path to store the booking state file")
	cobra.OnInitialize(func() {
		rootBookingState = state.NewFileState(statePath)
	})
	rootCmd.AddCommand(NewBookingCommand())
	rootCmd.AddCommand(NewCancelCommand())

	return rootCmd
}
