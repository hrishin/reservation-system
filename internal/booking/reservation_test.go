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
	want := true

	got, err := fr.BookSeats("A", 0, 2)
	if got != want || err != nil {
		t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
		t.Errorf("booking error: %v\n", err)
	}
}

func TestBookSeat_error_when_seats_are_beyond_capacity(t *testing.T) {
	temp := t.TempDir()
	st := state.NewFileState(temp)
	fr := NewFlightReservations(st)
	invalidSeats := 9
	want := false

	got, err := fr.BookSeats("A", 0, invalidSeats)
	if got != want || err == nil {
		t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
		t.Errorf("booking error: %v\n", err)
	}
}

func TestBookSeate_error_when_seats_are_booked_already(t *testing.T) {
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
	startSeats := 3
	invalidSeats := 3
	want := false

	got, err := fr.BookSeats("A", startSeats, invalidSeats)
	if got != want || err == nil {
		t.Errorf("failed to book the tikcet, expecting %v, got %v\n", want, got)
		t.Errorf("booking error: %v\n", err)
	}
}
