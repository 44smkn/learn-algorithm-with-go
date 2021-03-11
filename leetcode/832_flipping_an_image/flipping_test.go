package flipping_test

import (
	flipping "impl-using-go/leetcode/832_flipping_an_image"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlipAndInvertImage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"test1",
			[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := flipping.FlipAndInvertImage(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("want: %#v, got: %#v", tt.want, got)
			}
		})
	}
}
