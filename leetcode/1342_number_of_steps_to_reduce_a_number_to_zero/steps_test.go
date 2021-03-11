package steps_test

import (
	steps "impl-using-go/leetcode/1342_number_of_steps_to_reduce_a_number_to_zero"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"case 1", 14, 6},
		{"case 2", 8, 4},
		{"case 3", 123, 12},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := steps.NumberOfSteps(tt.input)
			if tt.want != got {
				t.Errorf("want: %v  got: %v\n", tt.want, got)
			}
		})
	}
}
