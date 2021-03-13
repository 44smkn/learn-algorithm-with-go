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
		{"case 1", []int{2, 2, 6, 6, 2, 3}, "4"},
		{"case 2", []int{2, 3, 3, 6, 2, 3, 8, 9}, "7"},
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
