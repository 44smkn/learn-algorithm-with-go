package findnums_test

import (
	findnums "impl-using-go/leetcode/1295_find_numbers_with_even_number_of_digits"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"case1", []int{12, 345, 2, 6, 7896}, 2},
		{"case2", []int{555, 901, 482, 1771}, 1},
		{"case3", []int{100000}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := findnums.FindNumbers(tt.input)
			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, tt.input)
			}
		})
	}
}
