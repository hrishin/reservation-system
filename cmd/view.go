package cmd

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/view"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

func NewViewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "VIEW",
		Short: "View flight booking",
		RunE: func(cmd *cobra.Command, args []string) error {
			bookings, err := rootBookingState.Load()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error in viewing flight bookings: %v\n", err)
				return err
			}

			seats := bookings.Seats
			var keys []string
			for key := range seats {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			// Iterate over the sorted keys and access the corresponding values from the map
			for _, row := range keys {
				seatData := seats[row]
				fmt.Printf("\t%s\n", view.FormatOccupancy(row, seatData))
			}

			return nil
		},
	}
}
