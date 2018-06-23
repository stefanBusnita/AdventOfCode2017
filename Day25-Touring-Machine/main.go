package main

import "fmt"

type Tape struct {
	content []int
}

func NewTape(size int) *Tape {
	return &Tape{
		content: make([]int, size),
	}
}

type State struct {
	config map[int]*Effect
	tape   *Tape
}

func NewState(tape *Tape) *State {
	cfg := make(map[int]*Effect)
	return &State{
		cfg,
		tape,
	}
}

// addEffect to the state of the machine
func (s *State) addEffect(val int, effect *Effect) {
	s.config[val] = effect
}

type Effect struct {
	nextStateKey string
}

func (Effect) doEffect(cursor int) {

}

func main() {
	tape := NewTape(1245)
	fmt.Println(tape)

}
