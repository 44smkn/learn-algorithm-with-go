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
		k    int
		s    int
		want string
	}{
		{"case 1", 2, 2, "6"},
		{"case 2", 5, 15, "1"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			input := bytes.NewBufferString(fmt.Sprintf("%v\n%v", tt.k, tt.s))
			buffer := &bytes.Buffer{}
			solve(input, buffer)
			got := strings.TrimSpace(buffer.String())
			if tt.want != got {
				t.Errorf("want: %v got: %v", tt.want, got)
			}
		})
	}
}
