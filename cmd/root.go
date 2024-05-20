package cmd

import (
	"github.com/hrishin/reservation-system/internal/booking"
	"github.com/hrishin/reservation-system/internal/state"
	"github.com/spf13/cobra"
	"log/slog"
	"math"
)

var statePath string

var rootBookingState booking.Storable

var verbose bool

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:           "booking",
		Short:         "A CLI application for booking and cancellation of flight reservations",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVar(&statePath, "state-file", "", "directory path to store the booking state file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enables verbose logging")

	cobra.OnInitialize(func() {
		rootBookingState = state.NewFileState(statePath)
		if !verbose {
			slog.SetLogLoggerLevel(math.MaxInt)
		}
	})
	rootCmd.AddCommand(NewBookingCommand())
	rootCmd.AddCommand(NewCancelCommand())
	rootCmd.AddCommand(NewViewCommand())

	return rootCmd
}
