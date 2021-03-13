package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a    []int
		want string
	}{
		{"case 1", []int{8, 12, 40}, "2"},
		{"case 2", []int{5, 6, 8, 10}, "0"},
		{"case 3", []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, "8"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fmtSlice := strings.Trim(fmt.Sprintf("%v", tt.a), "[]")
			input := bytes.NewBufferString(fmt.Sprintf("%v\n%v", len(tt.a), fmtSlice))
			buffer := &bytes.Buffer{}
			solve(input, buffer)
			got := strings.TrimSpace(buffer.String())
			if tt.want != got {
				t.Errorf("want: %v got: %v", tt.want, got)
			}
		})
	}
}
