package booking

import (
	"encoding/json"
	"github.com/hrishin/reservation-system/internal/state"
	"testing"
)

func TestBookSeat_success(t *testing.T) {
	temp := t.TempDir()
	st := state.NewFileState(temp)
	fr := NewFlightReservations(st)
	want := Success

	got, err := fr.BookSeats("A", 0, 2)
	if got != want || err != nil {
		t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
		t.Errorf("booking error: %v\n", err)
	}
}

func TestBookSeat_error_when_invalid_booking_request(t *testing.T) {
	tempDir := t.TempDir()
	seatsData := `{
		"seats": {
			"A": [1, 1, 1, -1, -1, 2, 2, 2],
			"B": [-1, -1, -1, -1, -1, -1, -1, -1],
			"C": [-1, -1, -1, -1, -1, -1, -1, -1],
			"D": [-1, -1, -1, -1, -1, -1, -1, -1],
			"E": [-1, -1, -1, -1, -1, -1, -1, -1],
			"F": [-1, -1, -1, -1, -1, -1, -1, -1],
			"G": [-1, -1, -1, -1, -1, -1, -1, -1],
			"H": [-1, -1, -1, -1, -1, -1, -1, -1],
			"I": [-1, -1, -1, -1, -1, -1, -1, -1],
			"J": [-1, -1, -1, -1, -1, -1, -1, -1],
			"K": [-1, -1, -1, -1, -1, -1, -1, -1],
			"L": [-1, -1, -1, -1, -1, -1, -1, -1],
			"M": [-1, -1, -1, -1, -1, -1, -1, -1],
			"N": [-1, -1, -1, -1, -1, -1, -1, -1],
			"O": [-1, -1, -1, -1, -1, -1, -1, -1],
			"P": [-1, -1, -1, -1, -1, -1, -1, -1],
			"Q": [-1, -1, -1, -1, -1, -1, -1, -1],
			"R": [-1, -1, -1, -1, -1, -1, -1, -1],
			"S": [-1, -1, -1, -1, -1, -1, -1, -1],
			"T": [-1, -1, -1, -1, -1, -1, -1, -1]
		},
		"id": 2
	}`
	var bookingState state.State
	err := json.Unmarshal([]byte(seatsData), &bookingState)
	if err != nil {
		t.Errorf("Error parsing JSON: %v\n", err)
	}
	storable := state.NewFileState(tempDir)
	err = storable.Save(&bookingState)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	st := state.NewFileState(tempDir)
	fr := NewFlightReservations(st)

	t.Run("seats exceeds beyond row capacity", func(t *testing.T) {
		want := Fail

		got, err := fr.BookSeats("A", 0, 9)
		if got != want {
			t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
			t.Errorf("booking error: %v\n", err)
		}
	})

	t.Run("some seats are already booked", func(t *testing.T) {
		want := Fail

		got, err := fr.BookSeats("A", 3, 3)
		if got != want {
			t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
			t.Errorf("booking error: %v\n", err)
		}
	})
}

func TestCancelBooking_success(t *testing.T) {
	tempDir := t.TempDir()
	seatsData := `{
		"seats": {
			"A": [1, 1, 1, -1, -1, 2, 2, 2],
			"B": [-1, -1, -1, -1, -1, -1, -1, -1],
			"C": [-1, -1, -1, -1, -1, -1, -1, -1],
			"D": [-1, -1, -1, -1, -1, -1, -1, -1],
			"E": [-1, -1, -1, -1, -1, -1, -1, -1],
			"F": [-1, -1, -1, -1, -1, -1, -1, -1],
			"G": [-1, -1, -1, -1, -1, -1, -1, -1],
			"H": [-1, -1, -1, -1, -1, -1, -1, -1],
			"I": [-1, -1, -1, -1, -1, -1, -1, -1],
			"J": [-1, -1, -1, -1, -1, -1, -1, -1],
			"K": [-1, -1, -1, -1, -1, -1, -1, -1],
			"L": [-1, -1, -1, -1, -1, -1, -1, -1],
			"M": [-1, -1, -1, -1, -1, -1, -1, -1],
			"N": [-1, -1, -1, -1, -1, -1, -1, -1],
			"O": [-1, -1, -1, -1, -1, -1, -1, -1],
			"P": [-1, -1, -1, -1, -1, -1, -1, -1],
			"Q": [-1, -1, -1, -1, -1, -1, -1, -1],
			"R": [-1, -1, -1, -1, -1, -1, -1, -1],
			"S": [-1, -1, -1, -1, -1, -1, -1, -1],
			"T": [3, 3, -1, -1, -1, -1, -1, -1]
		},
		"id": 4
	}`
	// Parse the JSON data into a State struct
	var bookingState state.State
	err := json.Unmarshal([]byte(seatsData), &bookingState)
	if err != nil {
		t.Errorf("Error parsing JSON: %v\n", err)
	}
	storable := state.NewFileState(tempDir)
	err = storable.Save(&bookingState)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	st := state.NewFileState(tempDir)
	fr := NewFlightReservations(st)
	want := Success

	got, err := fr.CancelSeats("T", 0, 2)
	if got != want {
		t.Errorf("failed to cancel the tikcet, expecting %v, got %v\n", want, got)
		t.Errorf("booking error: %v\n", err)
	}
}

func TestCancelBooking_error_when_invalid_cancel_request(t *testing.T) {
	tempDir := t.TempDir()
	seatsData := `{
		"seats": {
			"A": [-1, -1, -1, -1, -1, 2, 2, 2],
			"B": [4, 4, 4, 4, 4, 4, 4, 4],
			"C": [-1, -1, -1, -1, -1, -1, -1, -1],
			"D": [-1, -1, -1, -1, -1, -1, -1, -1],
			"E": [-1, -1, -1, -1, -1, -1, -1, -1],
			"F": [-1, -1, -1, -1, -1, -1, -1, -1],
			"G": [-1, -1, -1, -1, -1, -1, -1, -1],
			"H": [-1, -1, -1, -1, -1, -1, -1, -1],
			"I": [-1, -1, -1, -1, -1, -1, -1, -1],
			"J": [-1, -1, -1, -1, -1, -1, -1, -1],
			"K": [-1, -1, -1, -1, -1, -1, -1, -1],
			"L": [-1, -1, -1, -1, -1, -1, -1, -1],
			"M": [-1, -1, -1, -1, -1, -1, -1, -1],
			"N": [-1, -1, -1, -1, -1, -1, -1, -1],
			"O": [-1, -1, -1, -1, -1, -1, -1, -1],
			"P": [-1, -1, -1, -1, -1, -1, -1, -1],
			"Q": [-1, -1, -1, -1, -1, -1, -1, -1],
			"R": [-1, -1, -1, -1, -1, -1, -1, -1],
			"S": [-1, -1, -1, -1, -1, -1, -1, -1],
			"T": [3, 3, -1, -1, -1, -1, -1, -1]
		},
		"id": 5
	}`
	// Parse the JSON data into a State struct
	var bookingState state.State
	err := json.Unmarshal([]byte(seatsData), &bookingState)
	if err != nil {
		t.Errorf("Error parsing JSON: %v\n", err)
	}
	storable := state.NewFileState(tempDir)
	err = storable.Save(&bookingState)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	st := state.NewFileState(tempDir)
	fr := NewFlightReservations(st)

	t.Run("seat is not part of the reservation", func(t *testing.T) {
		want := Fail

		got, err := fr.CancelSeats("T", 0, 3)
		if got != want {
			t.Errorf("failed to cancel the tikcet, expecting %v, got %v\n", want, got)
		}
		if err == nil {
			t.Errorf("expecting error got %v\n", err)
		}
	})

	t.Run("seat has been never booked", func(t *testing.T) {
		want := Fail

		got, err := fr.CancelSeats("A", 0, 3)
		if got != want {
			t.Errorf("failed to cancel the tikcet, expecting %v, got %v\n", want, got)
		}
		if err == nil {
			t.Errorf("expecting error got %v\n", err)
		}
	})

	t.Run("cancellation request exceeds the row limit", func(t *testing.T) {
		want := Fail

		got, err := fr.CancelSeats("B", 0, 9)
		if got != want {
			t.Errorf("failed to cancel the tikcet, expecting %v, got %v\n", want, got)
		}
		if err == nil {
			t.Errorf("expecting error got %v\n", err)
		}
	})
}
