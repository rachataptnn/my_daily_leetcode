package main

import (
	"testing"
)

func TestNumRollsToTarget(t *testing.T) {
	type testCaseStruct struct {
		n        int
		k        int
		target   int
		expected int
	}
	testCases := []testCaseStruct{
		{
			n:        1,
			k:        6,
			target:   3,
			expected: 1,
		},
		{
			n:        2,
			k:        6,
			target:   7,
			expected: 6,
		},
		{
			n:        30,
			k:        30,
			target:   500,
			expected: 222616187,
		},
	}
	for _, testCase := range testCases {
		actual := numRollsToTarget(testCase.n, testCase.k, testCase.target)
		if actual != testCase.expected {
			t.Errorf(`
			n: %d
			k: %d
			target: %d
			
			expect: %d
			butgot: %d`, testCase.n, testCase.k, testCase.target, testCase.expected, actual)
		}
	}
}
