package view

import (
	"testing"
)

func TestFormatView(t *testing.T) {
	want := `A 11_1XX2_22`
	got := FormatOccupancy("A", []int{1, 1, 1, -1, -1, 2, 2, 2})
	if want != got {
		t.Errorf("got %v, want %v\n", got, want)
	}
}
