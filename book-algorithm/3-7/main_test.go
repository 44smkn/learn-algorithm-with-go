// https://atcoder.jp/contests/abc045/tasks/arc061_a?lang=ja
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "case1", s: "125", want: "176"},
		{name: "case2", s: "9999999999", want: "12656242944"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			input := bytes.NewBufferString(tt.s)
			buffer := &bytes.Buffer{}
			solve(input, buffer)
			got := strings.TrimSpace(buffer.String())
			if got != tt.want {
				t.Errorf("want: %v\ngot: %v", tt.want, got)
			}
		})
	}
}
