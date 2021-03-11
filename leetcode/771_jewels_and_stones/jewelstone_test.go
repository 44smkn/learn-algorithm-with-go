package jewelstone_test

import (
	jewelstone "impl-using-go/leetcode/771_jewels_and_stones"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{"case 1", "aA", "aAAbbbb", 3},
		{"case 2", "z", "ZZ", 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := jewelstone.NumJewelsInStones(tt.jewels, tt.stones)
			if tt.want != got {
				t.Errorf("want: %v  got: %v", tt.want, got)
			}
		})
	}
}
