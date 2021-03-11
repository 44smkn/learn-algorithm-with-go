package datastructure

import (
	"errors"
	"strconv"
	"strings"
)

func calcWithStack(expression string) (int, error) {

	s := newStack(make([]int, 10))
	for _, v := range strings.Split(expression, " ") {
		switch v {
		case "+":
			a, err := s.pop()
			if err != nil {
				return 0, err
			}
			b, err := s.pop()
			if err != nil {
				return 0, err
			}
			err = s.push(b + a)
			if err != nil {
				return 0, err
			}
		case "-":
			a, err := s.pop()
			if err != nil {
				return 0, err
			}
			b, err := s.pop()
			if err != nil {
				return 0, err
			}
			err = s.push(b - a)
			if err != nil {
				return 0, err
			}
		case "*":
			a, err := s.pop()
			if err != nil {
				return 0, err
			}
			b, err := s.pop()
			if err != nil {
				return 0, err
			}
			err = s.push(b * a)
			if err != nil {
				return 0, err
			}
		default:
			num, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			err = s.push(num)
			if err != nil {
				return 0, err
			}
		}
	}

	return s.slice[s.top], nil
}

type stack struct {
	top   int
	max   int
	slice []int
}

func newStack(slice []int) *stack {
	return &stack{
		top:   0,
		max:   len(slice),
		slice: slice,
	}
}

func (s *stack) isEmpty() bool {
	return s.top == 0
}

func (s *stack) isFull() bool {
	return s.top == s.max-1
}

func (s *stack) push(x int) error {
	if s.isFull() {
		return errors.New("Overflow")
	}
	s.top++
	s.slice[s.top] = x
	return nil
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Underflow")
	}
	s.top--
	return s.slice[s.top+1], nil
}
