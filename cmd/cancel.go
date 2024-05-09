package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/booking"
	"log/slog"
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
				slog.Error(fmt.Sprintf("parsing seat preference, starting seat: %v\n", err))
				return err
			}

			seats, err := strconv.Atoi(numSeats)
			if err != nil {
				slog.Error(fmt.Sprintf("parsing seat preference, number of seat: %v\n", err))
				return err
			}

			row := string(seatPreference[0])

			reservation := booking.NewFlightReservations(rootBookingState)
			result, err := reservation.CancelSeats(row, start, seats)
			fmt.Println(result)
			return err
		},
	}
}
