package sum_test

import (
	sum "impl-using-go/leetcode/1480_running_sum"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRunningSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"4 elements", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"1 * 5 elements", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"5 elements", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sum.RunningSum(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
