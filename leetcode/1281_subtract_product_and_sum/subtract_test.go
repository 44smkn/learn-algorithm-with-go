package subtract_test

import (
	subtract "impl-using-go/leetcode/1281_subtract_product_and_sum"
	"testing"
)

func TestSubtractProductAndSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"case1", 234, 15},
		{"case2", 4421, 21},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := subtract.SubtractProductAndSum(tt.n)
			if tt.want != got {
				t.Errorf("want: %v \ngot: %v", tt.want, got)
			}
		})
	}
}
