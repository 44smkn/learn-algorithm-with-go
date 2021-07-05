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
		{name: "case1", n: "575", want: "4"},
		{name: "case2", n: "3600", want: "13"},
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
