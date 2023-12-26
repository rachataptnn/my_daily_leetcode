package main

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	type testCaseStruct struct {
		input    int
		expected []int
	}

	testCases := []testCaseStruct{
		{
			input:    2,
			expected: []int{0, 1, 1},
		},
		{
			input:    5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tc := range testCases {
		actual := countBits(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf(`
				input:    %v
				expected: %v
				but got:  %v`, tc.input, actual, tc.expected)
		}
	}
}
