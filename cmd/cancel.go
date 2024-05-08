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
		Use:   "CANCEL flight_id] [num_tickets]\n\n  CANCEL A0 1",
		Short: "Cancel a booking",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			seatPreference := args[0]
			numSeats := args[1]

			startNum := seatPreference[1:]
			start, err := strconv.Atoi(startNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting substring to integer: %v\n", err)
				return err
			}

			seats, err := strconv.Atoi(numSeats)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting substring to integer: %v\n", err)
				return err
			}

			row := string(seatPreference[0])

			reservation := booking.NewFlightReservations(rootBookingState)
			done, err := reservation.CancelSeats(row, start, seats)
			if !done {
				fmt.Fprintf(os.Stderr, "cancellation failed for %s tickets for seat %s : %v\n", numSeats, seatPreference, err)
				return err
			}
			fmt.Printf("Confirmed cancellation %s tickets for seating %s\n", numSeats, seatPreference)
			return nil
		},
	}
}
