package state

const EmptySeat = -1

type State struct {
	Seats map[string][]int `json:"seats"`
	ID    int              `json:"id"`
}
