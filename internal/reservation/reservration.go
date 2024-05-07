package reservation

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/state"
)

type FlightReservation struct {
	state state.Storable
}

func (r *FlightReservation) BookSeats(row string, start int, seats int) (bool, error) {
	rs, err := r.state.Load()
	if err != nil {
		fmt.Printf("error loading the reservation state: %v\n", err)
		return false, err
	}

	lastSeat := start + seats - 1
	maxSeat := len(rs.Seats[row])
	if _, ok := rs.Seats[row]; !ok || lastSeat >= maxSeat {
		return false, fmt.Errorf("number of seats are beyong the current capacity of the reservation")
	}

	for i := 0; i < seats; i++ {
		if rs.Seats[row][start+i] >= 0 {
			return false, fmt.Errorf("requested seat number request is already booked: %v%v", row, start+i)
		}
	}

	currentID := rs.ID
	for i := 0; i < seats; i++ {
		rs.Seats[row][start+i] = currentID
	}

	rs.ID = currentID + 1
	if err := r.state.Save(rs); err != nil {
		fmt.Printf("error in saving the reservation state: %v\n", err)
		return false, err
	}

	return true, nil
}

func NewFlightReservations(state state.Storable) *FlightReservation {
	return &FlightReservation{
		state: state,
	}
}
