package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, test := range tests {
		if got := fib(test.n); got != test.want {
			t.Errorf(`n: %d
			but got: %d
			`, test.n, got)
		}
	}
}
