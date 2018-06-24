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

const (
	ZERO WriteVal = 0
	ONE  WriteVal = 1
)

type Effect struct {
	nextStateKey string
	effect       func(Cursor) (Cursor, WriteVal)
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
	fmt.Printf("Cursor start value %d \n", cursor)

	iterations := 0
	var writeVal WriteVal

	for iterations < steps {

		state := t.states[startStateKey]
		currentTapeValue := t.content[cursor]

		// call do effect which should modify the tape
		startStateKey = state.config[currentTapeValue].nextStateKey

		cursor, writeVal = state.config[currentTapeValue].effect(cursor)
		//fmt.Printf("Current cursor %d with the next value %d \n", cursor, int(writeVal))
		t.content[int(cursor)] = int(writeVal)

		// just for testing, the sum of all these ones is the checksum
		if t.content[int(cursor)] == 1 {
			fmt.Println("One One")
		}
		//fmt.Printf("How does the tape look now ? %+v \n", t.content)

		iterations++
	}

}

func main() {
	tape := NewTape(50)

	// create and describe stateA
	stateA := NewState(tape)
	stateA.addEffect(0, &Effect{
		nextStateKey: "B",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})
	stateA.addEffect(1, &Effect{
		nextStateKey: "B",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ZERO
		},
	})
	// create and describe stateB
	stateB := NewState(tape)
	stateB.addEffect(0, &Effect{
		nextStateKey: "A",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ONE
		},
	})
	stateB.addEffect(1, &Effect{
		nextStateKey: "A",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})

	tape.addState("A", stateA)
	tape.addState("B", stateB)

	tape.start(5, "A")

}
