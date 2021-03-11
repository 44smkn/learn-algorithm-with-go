package shuffle_test

import (
	shuffle "impl-using-go/leetcode/1470_shuffle_the_array"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShuffle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		// input
		nums []int
		n    int
		// output
		want []int
	}{
		{"6 elements", []int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{"8 elements", []int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"4 elements", []int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, tt := range tests {
		tt := tt // シャドウイングしないとRunが走るタイミングが逐次ではないため、最後の要素が3回行われる
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := shuffle.Shuffle(tt.nums, tt.n)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
