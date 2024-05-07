package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const STATE_FILE_NAME = "seat_sate"

type ReservationStateFile struct {
	path string
	file string
}

func (r *ReservationStateFile) Load() (*State, error) {
	//TODO implement me
	var seatState State
	sf := filepath.Join(r.path, r.file)
	file, err := os.Open(sf)

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error loading seat state:", err)
			return nil, err
		}

		seatState = State{ID: 1}
		for i := 'A'; i <= 'U'; i++ {
			seatState.seats[string(i)] = make([]int32, 8)
			for j := range seatState.seats[string(i)] {
				seatState.seats[string(i)][j] = -1
			}
		}
		return &seatState, nil
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&seatState)
	if err != nil {
		fmt.Println("Error decoding seat state:", err)
		return nil, err
	}

	return &seatState, err
}

func (r *ReservationStateFile) Save(seatState *State) error {
	sf := filepath.Join(r.path, r.file)
	file, err := os.Create(sf)
	if err != nil {
		fmt.Println("Error saving seat state:", err)
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(seatState)
	if err != nil {
		fmt.Println("Error encoding seat state:", err)
		return err
	}
	return nil
}

func NewFileState(sateDir string) Storable {
	if sateDir == "" {
		// Get the user's home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			//TODO: re-think exit
			fmt.Println("Error finding user's home directory for storing the default sate file:", err)
			os.Exit(0)
		}
		sateDir = filepath.Join(homeDir, ".reservation")
	}

	return &ReservationStateFile{
		path: sateDir,
		file: STATE_FILE_NAME,
	}
}