package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/booking"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func NewBookingCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "BOOK [flight_id] [num_tickets]\n\n  BOOK A0 1",
		Short: "Book a flight",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			seatPreference := args[0]
			numSeats := args[1]

			startNum := seatPreference[1:]
			start, err := strconv.Atoi(startNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error converting substring to integer: %v\n", err)
				return err
			}

			seats, err := strconv.Atoi(numSeats)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error converting substring to integer: %v\n", err)
				return err
			}

			reservation := booking.NewFlightReservations(rootBookingState)
			done, err := reservation.BookSeats(string(seatPreference[0]), start, seats)
			if !done {
				fmt.Fprintf(os.Stderr, "booking failed for %s tickets for seat %s : %v\n", numSeats, seatPreference, err)
				return err
			}
			fmt.Printf("confirmed %s tickets for seating %s\n", numSeats, seatPreference)
			return nil
		},
	}
}
