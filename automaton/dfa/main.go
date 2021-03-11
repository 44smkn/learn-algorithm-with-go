package main

import (
	"fmt"
)

// State is represented with ○ in figure.
type State int

// DeterministicFiniteAutomaton is M=(Q,Σ,δ,q0,F)
type DeterministicFiniteAutomaton struct {

	// Q = set of state.
	states map[State]bool

	// Σ = alphabet. It is set of character. But it is not only [a-z][A-Z]{0,}.
	// For example,
	// Σ = {0, 1}
	// Σ = {0, 1, 2, ..., 9, a,b,c, ..., z, A, B, C, ..., Z, #, $, %}
	alphabets map[rune]bool

	// transition function
	transitions func(state State, alphabet rune) State

	// q0 = initial state
	initState State

	// F = set of acceptence state 
	finalState State
}

func (d *DeterministicFiniteAutomaton) judge(input string) bool {
	currentState := d.initState
	for _, v := range input {
		if _, ok := d.alphabets[v]; !ok {
			continue
		}
		currentState = d.transitions(currentState, v)
	}
	return currentState == d.finalState
}

func main() {

	q := newStateSet(0, 1, 2)
	alphabets := newAlphabetSet('0', '1', 'a')
	transitions := func(state State, alphabet rune) State {
		m := map[State]map[rune]State {
			0: {
				'0': 0,
				'1': 0,
				'a': 1,
			},
			1: {
				'0': 1,
				'1': 1,
				'a': 2,
			},
			2: {
				'0': 2,
				'1': 2,
				'a': 2,
			},
		}
		return m[state][alphabet]
	}

	dfa := DeterministicFiniteAutomaton {
		states: q,
		alphabets: alphabets,
		transitions: transitions,
		initState: 0,
		finalState: 2,
	}

	inputs := []string{
		"01a0a",
		"1001a",
		"aa020",
		"a1001",
	}

	for _, v := range inputs {
		fmt.Printf("input: %v, accept: %v\n", v, dfa.judge(v))
	}
}

func newStateSet(elements ...State) map[State]bool {
	set := make(map[State]bool)
	for _, e := range elements {
		set[e] = true
	}
	return set
}

func newAlphabetSet(elements ...rune)  map[rune]bool {
	set := make(map[rune]bool)
	for _, e := range elements {
		set[e] = true
	}
	return set
}