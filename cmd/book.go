package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/state"
	"github.com/spf13/cobra"
)

func NewBookingCommand(state.Storable) *cobra.Command {
	return &cobra.Command{
		Use:   "BOOK [flight_id] [num_tickets]",
		Short: "Book a flight",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			seatPreference := args[0]
			numSeats := args[1]
			fmt.Printf("Booking %s tickets for flight %s\n", numSeats, seatPreference)
			// Your booking logic here
		},
	}
}
