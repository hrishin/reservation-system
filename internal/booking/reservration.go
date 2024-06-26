package booking

import (
	"fmt"
	"github.com/hrishin/reservation-system/internal/state"
	"log/slog"
)

type Result string

const (
	Success Result = "SUCCESS"
	Fail    Result = "FAIL"
)

type Storable interface {
	Load() (*state.State, error)
	Save(state *state.State) error
}

type FlightReservation struct {
	state Storable
}

func (r *FlightReservation) BookSeats(row string, start int, seats int) (Result, error) {
	rs, err := r.state.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("loading the booking state: %v\n", err))
		return Fail, err
	}

	lastSeat := start + seats - 1
	maxSeat := len(rs.Seats[row])
	if _, ok := rs.Seats[row]; !ok || isBeyondSeatsCapacity(lastSeat, maxSeat) {
		return Fail, fmt.Errorf("number of seats are beyond the current capacity of the booking")
	}

	for i := 0; i < seats; i++ {
		if isSeatBooked(row, start+i, rs) {
			return Fail, fmt.Errorf("requested seat number request is already booked: %v%v", row, start+i)
		}
	}

	currentID := rs.ID
	for i := 0; i < seats; i++ {
		rs.Seats[row][start+i] = currentID
	}

	rs.ID = currentID + 1
	if err := r.state.Save(rs); err != nil {
		slog.Error(fmt.Sprintf("updating the booking state: %v\n", err))
		return Fail, err
	}

	return Success, nil
}

func (r *FlightReservation) CancelSeats(row string, start int, seats int) (Result, error) {
	rs, err := r.state.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("loading the booking state: %v\n", err))
		return Fail, err
	}

	lastSeat := start + seats - 1
	maxSeat := len(rs.Seats[row])
	if _, ok := rs.Seats[row]; !ok || isBeyondSeatsCapacity(lastSeat, maxSeat) {
		return Fail, fmt.Errorf("number of seats are beyond the current capacity of the booking or sea prference is invalid")
	}

	id := rs.Seats[row][start]
	for i := 0; i < seats; i++ {
		if !isFromSameBookingRequest(id, row, start+i, rs) || !isSeatBooked(row, start+i, rs) {
			return Fail, fmt.Errorf("requested seat %v%v is not valid for cancellation, either not booked for this booking or never booked before at all", row, start+i)
		}
	}

	for i := 0; i < seats; i++ {
		rs.Seats[row][start+i] = state.EmptySeat
	}

	if err := r.state.Save(rs); err != nil {
		slog.Error(fmt.Sprintf("updating the booking state: %v\n", err))
		return Fail, err
	}

	return Success, nil
}

func isBeyondSeatsCapacity(seat int, maxSeat int) bool {
	return seat >= maxSeat
}

func isSeatBooked(row string, seat int, rs *state.State) bool {
	return rs.Seats[row][seat] >= 0
}

func isFromSameBookingRequest(id int, row string, seat int, rs *state.State) bool {
	return rs.Seats[row][seat] == id
}

func NewFlightReservations(state Storable) *FlightReservation {
	return &FlightReservation{
		state: state,
	}
}
