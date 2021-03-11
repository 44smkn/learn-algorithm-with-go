package smaller_test

import (
	smaller "impl-using-go/leetcode/1365_smaller_numbers_than_current"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"case1", []int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{"case2", []int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{"case3", []int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := smaller.SmallerNumbersThanCurrent(tt.nums)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
