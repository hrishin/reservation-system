package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/booking"
	"github.com/hrishin/reservation-system/internal/state"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func NewBookingCommand(bookingState state.Storable) *cobra.Command {
	return &cobra.Command{
		Use:   "BOOK [flight_id] [num_tickets]",
		Short: "Book a flight",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			seatPreference := args[0]
			numSeats := args[1]

			startNum := seatPreference[1:]
			start, err := strconv.Atoi(startNum)
			if err != nil {
				fmt.Printf("Error converting substring to integer: %v\n", err)
				return
			}

			seats, err := strconv.Atoi(numSeats)
			if err != nil {
				fmt.Printf("Error converting substring to integer: %v\n", err)
				return
			}

			reservoir := booking.NewFlightReservations(bookingState)
			done, err := reservoir.BookSeats(string(seatPreference[0]), start, seats)
			if !done {
				fmt.Printf("booking failed for %s tickets for seat %s : %v\n", numSeats, seatPreference, err)
				os.Exit(-1)
			}
			fmt.Printf("Confirmed %s tickets for seating %s\n", numSeats, seatPreference)
		},
	}
}
