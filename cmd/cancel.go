package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/state"

	"github.com/spf13/cobra"
)

func NewCancelCommand(state.Storable) *cobra.Command {
	return &cobra.Command{
		Use:   "CANCE [booking_id]",
		Short: "Cancel a booking",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			bookingID := args[0]
			fmt.Printf("Cancelling booking %s\n", bookingID)
			// Your cancellation logic here
		},
	}
}
