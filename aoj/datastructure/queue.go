package datastructure

import (
	"errors"
	"sort"
)

type process struct {
	name string
	time int
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func proceedWithQueue(n int, quantum int, times map[string]int) (map[string]int, error) {
	keys := make([]string, 0, len(times))
	for name := range times {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	processes := make([]process, n)
	for i, name := range keys {
		time := times[name]
		proc := process{
			name: name,
			time: time,
		}
		processes[i] = proc
	}

	q := newQueue(make([]process, 8))
	for _, v := range processes {
		err := q.enqueue(&v)
		if err != nil {
			return nil, err
		}
	}

	var elaps int
	output := make(map[string]int, n)
	for !q.isEmpty() {
		p, err := q.dequeue()
		if err != nil {
			return nil, err
		}
		c := min(quantum, p.time)
		p.time -= c
		elaps += c
		if p.time > 0 {
			err := q.enqueue(p)
			if err != nil {
				return nil, err
			}
		} else {
			output[p.name] = elaps
		}
	}
	return output, nil
}

type queue struct {
	head  int
	tail  int
	slice []process
	max   int
}

func newQueue(slice []process) *queue {
	return &queue{
		head:  0,
		tail:  0,
		max:   len(slice),
		slice: slice,
	}
}

func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

// On array of Ring Buffer, head is next to tail
func (q *queue) isFull() bool {
	return q.head == (q.tail+1)%q.max
}

func (q *queue) enqueue(x *process) error {
	if q.isFull() {
		return errors.New("Overflow")
	}
	q.slice[q.tail] = *x
	if q.tail == q.max-1 {
		q.tail = 0
	} else {
		q.tail++
	}
	return nil
}

func (q *queue) dequeue() (*process, error) {
	if q.isEmpty() {
		return nil, errors.New("Underflow")
	}
	x := &q.slice[q.head]
	if q.head == q.max-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x, nil
}
