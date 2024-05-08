package view

import (
	"strconv"
	"strings"
)

// formatOccupancy formats the seat occupancy slice into a concise string format
func FormatOccupancy(row string, seats []int) string {
	var builder strings.Builder
	builder.WriteString(row)
	builder.WriteString(" ")
	for i, seat := range seats {
		if int(seat) > 0 {
			builder.WriteString(strconv.Itoa(seat))
		} else {
			builder.WriteString("X")
		}

		// Add separator after every two seats except the last seat
		if i == 1 || i == 5 {
			builder.WriteString("_")
		}
	}

	return builder.String()
}
