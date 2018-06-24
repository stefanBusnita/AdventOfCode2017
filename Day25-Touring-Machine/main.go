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
		content: make([]int, size*10),
		states:  make(map[string]*State, 0),
	}
}

func (t *Tape) addState(key string, state *State) {
	t.states[key] = state
}

func (t *Tape) start(steps int, key string) {

	cursor := Cursor(steps * 10 / 2)
	fmt.Printf("Cursor start value %d and state %s \n", cursor, key)

	iterations := 0
	var writeVal WriteVal
	var newCursor Cursor

	for iterations < steps {

		state := t.states[key]
		currentTapeValue := t.content[cursor]

		fmt.Printf("State %s, Cursor %d,current value %d, next value %d \n", key, cursor, currentTapeValue, int(writeVal))
		newCursor, writeVal = state.config[currentTapeValue].effect(cursor)

		fmt.Printf("State %s, Cursor %d,current value %d, next value %d \n \n\n\n", key, cursor, currentTapeValue, int(writeVal))
		// assign new value
		t.content[int(cursor)] = int(writeVal)
		// get the next state key
		key = state.config[currentTapeValue].nextStateKey
		// move cursor ahead
		cursor = newCursor
		// count the next iteration
		iterations++
	}
	onez2 := 0
	for i := 0; i < len(t.content); i++ {
		if t.content[i] == 1 {
			onez2++
		}
	}

	fmt.Printf("Final count %d \n", onez2)

}

func main() {
	tape := NewTape(12794428)

	/*	In state A:
		If the current value is 0:
		  - Write the value 1.
		  - Move one slot to the right.
		  - Continue with state B.
		If the current value is 1:
		  - Write the value 0.
		  - Move one slot to the left.
		  - Continue with state F.*/

	// create and describe stateA
	stateA := NewState(tape)
	stateA.addEffect(0, &Effect{
		nextStateKey: "B",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})
	stateA.addEffect(1, &Effect{
		nextStateKey: "F",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ZERO
		},
	})

	/*
			In state B:
		  If the current value is 0:
		    - Write the value 0.
		    - Move one slot to the right.
		    - Continue with state C.
		  If the current value is 1:
		    - Write the value 0.
		    - Move one slot to the right.
		    - Continue with state D.
	*/
	// create and describe stateB
	stateB := NewState(tape)
	stateB.addEffect(0, &Effect{
		nextStateKey: "C",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ZERO
		},
	})
	stateB.addEffect(1, &Effect{
		nextStateKey: "D",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ZERO
		},
	})

	/*In state C:
	  If the current value is 0:
	    - Write the value 1.
	    - Move one slot to the left.
	    - Continue with state D.
	  If the current value is 1:
	    - Write the value 1.
	    - Move one slot to the right.
	    - Continue with state E.*/
	stateC := NewState(tape)
	stateC.addEffect(0, &Effect{
		nextStateKey: "D",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ONE
		},
	})
	stateC.addEffect(1, &Effect{
		nextStateKey: "E",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})
	/*

		In state D:
		  If the current value is 0:
		    - Write the value 0.
		    - Move one slot to the left.
		    - Continue with state E.
		  If the current value is 1:
		    - Write the value 0.
		    - Move one slot to the left.
		    - Continue with state D.
	*/
	stateD := NewState(tape)
	stateD.addEffect(0, &Effect{
		nextStateKey: "E",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ZERO
		},
	})
	stateD.addEffect(1, &Effect{
		nextStateKey: "D",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ZERO
		},
	})

	/*

		In state E:
		  If the current value is 0:
		    - Write the value 0.
		    - Move one slot to the right.
		    - Continue with state A.
		  If the current value is 1:
		    - Write the value 1.
		    - Move one slot to the right.
		    - Continue with state C.
	*/

	stateE := NewState(tape)
	stateE.addEffect(0, &Effect{
		nextStateKey: "A",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ZERO
		},
	})
	stateE.addEffect(1, &Effect{
		nextStateKey: "C",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})

	/*
		In state F:
		  If the current value is 0:
		    - Write the value 1.
		    - Move one slot to the left.
		    - Continue with state A.
		  If the current value is 1:
		    - Write the value 1.
		    - Move one slot to the right.
			- Continue with state A.

	*/

	stateF := NewState(tape)
	stateF.addEffect(0, &Effect{
		nextStateKey: "A",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor - 1, ONE
		},
	})
	stateF.addEffect(1, &Effect{
		nextStateKey: "A",
		effect: func(cursor Cursor) (Cursor, WriteVal) {
			return cursor + 1, ONE
		},
	})

	tape.addState("A", stateA)
	tape.addState("B", stateB)
	tape.addState("C", stateC)
	tape.addState("D", stateD)
	tape.addState("E", stateE)
	tape.addState("F", stateF)

	tape.start(12794428, "A")

	// 3198607 Wrong !

	// 3198607 MIGHT BE GOOD, WE CAN TRY IN ONE MINUTE

}
