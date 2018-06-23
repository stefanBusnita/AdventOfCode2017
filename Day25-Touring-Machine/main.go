package main

import "fmt"

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

type Cursor int
type WriteVal int

type Effect struct {
	nextStateKey string
	effect       func(cursor int) (Cursor, WriteVal)
}

func doEffect(cursor int) (Cursor, WriteVal) {
	return Cursor(cursor + 1), WriteVal(1)
}

func doEffect2(cursor int) (Cursor, WriteVal) {
	return Cursor(cursor - 1), WriteVal(0)
}

func doEffect3(cursor int) (Cursor, WriteVal) {
	return Cursor(cursor - 1), WriteVal(1)
}

func doEffect4(cursor int) (Cursor, WriteVal) {
	return Cursor(cursor + 1), WriteVal(1)
}

type Tape struct {
	content []int
	states  map[string]*State
}

func NewTape(size int) *Tape {
	return &Tape{
		content: make([]int, size),
		states:  make(map[string]*State, 0),
	}
}

func (t *Tape) addState(key string, state *State) {
	t.states[key] = state
}

func (t *Tape) start(steps int, startStateKey string) {

	cursor := Cursor(steps * 10 / 2)
	fmt.Printf("Cursor start value %d", cursor)

	iterations := 0
	var writeVal WriteVal

	for iterations < steps {

		state := t.states[startStateKey]
		currentTapeValue := t.content[cursor]

		// call do effect which should modify the tape
		startStateKey = state.config[currentTapeValue].nextStateKey

		cursor, writeVal = state.config[currentTapeValue].effect(int(cursor))
		fmt.Printf("Current cursor %d with the next value %d \n", cursor, int(writeVal))
		t.content[int(cursor)] = int(writeVal)

		// just for testing, the sum of all these ones is the checksum
		if t.content[int(cursor)] == 1 {
			fmt.Println("One One")
		}

		iterations++
	}

}

func main() {
	tape := NewTape(200)

	// create and describe stateA
	stateA := NewState(tape)
	stateA.addEffect(0, &Effect{
		nextStateKey: "B",
		effect:       doEffect,
	})
	stateA.addEffect(1, &Effect{
		nextStateKey: "B",
		effect:       doEffect2,
	})
	// create and describe stateB
	stateB := NewState(tape)
	stateB.addEffect(0, &Effect{
		nextStateKey: "A",
		effect:       doEffect3,
	})
	stateB.addEffect(1, &Effect{
		nextStateKey: "A",
		effect:       doEffect4,
	})

	tape.addState("A", stateA)
	tape.addState("B", stateB)

	tape.start(5, "A")

	fmt.Println(tape)
}
