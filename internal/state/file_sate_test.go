package state

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFileStatLoadWhenStateFileDoesntExist(t *testing.T) {
	// Test with non-empty stateDir argument
	tempDir := t.TempDir() // Create a temporary directory for testing
	// Define the JSON data as a string
	jsonData := `{
		"seats": {
			"A": [-1, -1, -1, -1, -1, -1, -1, -1],
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
		"id": 1
	}`
	// Parse the JSON data into a State struct
	var want State
	err := json.Unmarshal([]byte(jsonData), &want)
	if err != nil {
		t.Errorf("Error parsing JSON: %v\n", err)
	}

	storable := NewFileState(tempDir)
	got, err := storable.Load()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)

	}
	if diff := cmp.Diff(*got, want); diff != "" {
		t.Errorf("Diff: %v\nExpected %v, got %v\n", diff, *got, want)
	}
}

func TestFileStatSaveAndLoad(t *testing.T) {
	// Test with non-empty stateDir argument
	tempDir := t.TempDir() // Create a temporary directory for testing
	// Define the JSON data as a string
	jsonData := `{
		"seats": {
			"A": [1, 1, -1, -1, -1, -1, -1, -1],
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
	var state State
	err := json.Unmarshal([]byte(jsonData), &state)
	if err != nil {
		t.Errorf("Error parsing JSON: %v\n", err)
	}
	storable := NewFileState(tempDir)
	err = storable.Save(&state)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	got, err := storable.Load()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	if diff := cmp.Diff(*got, state); diff != "" {
		t.Errorf("Expected %v, got %v\n", state, *got)
	}
}

func TestNewFileState(t *testing.T) {
	// Test with empty stateDir argument
	storable := NewFileState("")
	if storable == nil {
		t.Error("Expected non-nil Storable, got nil")
	}

	// Test with non-empty stateDir argument
	tempDir := t.TempDir() // Create a temporary directory for testing

	storable = NewFileState(tempDir)
	if storable == nil {
		t.Error("Expected non-nil Storable, got nil")
	}

	// Test if the path is set correctly
	fileStat, ok := storable.(*ReservationStateFile)
	if !ok {
		t.Error("Unexpected type returned from NewFileState")
	}
	expectedPath := tempDir
	if fileStat.path != expectedPath {
		t.Errorf("Expected path %s, got %s", expectedPath, fileStat.path) // Ensure path is set correctly
	}
}
