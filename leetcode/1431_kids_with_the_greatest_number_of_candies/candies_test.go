package candies_test

import (
	"fmt"
	candies "impl-using-go/leetcode/1431_kids_with_the_greatest_number_of_candies"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKidsWithCandies(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		// input
		candies      []int
		extraCandies int

		want []bool
	}{
		{"5 elements, extra 3", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{"5 elements, extra 1", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{"3 elements, extra 10", []int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(*testing.T) {
			t.Parallel()
			got := candies.KidsWithCandies(tt.candies, tt.extraCandies)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf(fmt.Sprintf("you wants %v. but, you got %v.", tt.want, got))
			}
		})
	}
}
