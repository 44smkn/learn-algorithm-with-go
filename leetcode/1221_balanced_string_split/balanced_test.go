package balanced_test

import (
	balanced "impl-using-go/leetcode/1221_balanced_string_split"
	"testing"
)

func TestBalancedStringSplit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"case1", "RLRRLLRLRL", 4},
		{"case2", "RLLLLRRRLR", 3},
		{"case3", "LLLLRRRR", 1},
		{"case4", "RLRRRLLRLL", 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := balanced.BalancedStringSplit(tt.input)
			if tt.want != got {
				t.Errorf("want: %v  got: %v", tt.want, got)
			}
		})
	}
}
