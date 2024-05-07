package state

type State struct {
	seats map[string][]int32 `json:"seats"`
	ID    int32              `json:"id"`
}

type Storable interface {
	Load() (*State, error)
	Save(state *State) error
}
