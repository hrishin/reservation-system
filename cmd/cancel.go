package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/booking"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCancelCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "CANCEL [booking_id]",
		Short: "Cancel a booking",
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

			row := string(seatPreference[0])

			reservoir := booking.NewFlightReservations(rootBookingState)
			done, err := reservoir.CancelSeats(row, start, seats)
			if !done {
				fmt.Printf("cancellation failed for %s tickets for seat %s : %v\n", numSeats, seatPreference, err)
				os.Exit(-1)
			}
			fmt.Printf("Confirmed cancellation %s tickets for seating %s\n", numSeats, seatPreference)
		},
	}
}
