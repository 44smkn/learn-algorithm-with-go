package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		n    string
		want string
	}{
		{name: "case1", n: "9", want: "24"},
		{name: "case2", n: "1", want: "0"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			input := bytes.NewBufferString(tt.n)
			buffer := &bytes.Buffer{}
			solve(input, buffer)
			got := strings.TrimSpace(buffer.String())
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
