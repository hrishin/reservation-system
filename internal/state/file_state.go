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
	var seatState State
	sf := filepath.Join(r.path, r.file)
	file, err := os.Open(sf)

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error loading seat state:", err)
			return nil, err
		}

		seatState = State{ID: 1, Seats: make(map[string][]int)}
		for i := 'A'; i <= 'U'; i++ {
			row := string(i)
			seatState.Seats[row] = make([]int, 8)
			for j := range seatState.Seats[row] {
				seatState.Seats[row][j] = EmptySeat
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
	err := os.MkdirAll(r.path, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return err
	}

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
		sateDir = filepath.Join(homeDir, ".booking")
	}

	return &ReservationStateFile{
		path: sateDir,
		file: STATE_FILE_NAME,
	}
}
