package divideconquer_test

import (
	"impl-using-go/aoj/divideconquer"
	"testing"
)

func TestExhaustiveSearch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		A    []int
		m    []int
		want string
	}{
		{
			name: "example1",
			A:    []int{1, 5, 7, 10, 21},
			m:    []int{2, 4, 17, 8},
			want: "nonoyesyes",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := divideconquer.ExhaustiveSearch(tt.A, tt.m); got != tt.want {
				t.Errorf("got: %v want: %v", got, tt.want)
			}
		})
	}
}
