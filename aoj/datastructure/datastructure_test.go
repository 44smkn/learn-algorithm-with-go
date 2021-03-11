package datastructure

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQueue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		n       int
		quantam int
		times   map[string]int
		want    map[string]int
	}{
		{"case1", 5, 100, map[string]int{"p1": 150, "p2": 80, "p3": 200, "p4": 350, "p5": 20}, map[string]int{"p1": 450, "p2": 180, "p3": 550, "p4": 800, "p5": 400}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, _ := proceedWithQueue(tt.n, tt.quantam, tt.times)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestStack(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"case1", "1 2 + 3 4 - *", -3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, _ := calcWithStack(tt.input)
			if tt.want != got {
				t.Errorf("want: %v  got: %v", tt.want, got)
			}
		})
	}
}

func TestLinkedList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		instruction []string
		want        []int
	}{
		{
			"case1",
			[]string{"insert 5", "insert 2", "insert 3", "insert 1", "delete 3", "insert 6", "delete 5"},
			[]int{6, 1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := printList(tt.instruction)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
