package main

import (
	testing "testing"
)

func TestNumDecodings(t *testing.T) {
	type testCaseStruct struct {
		caseName string
		input    string
		expected int
	}

	testCases := []testCaseStruct{
		{"case 1", "12", 2},
		{"case 2", "226", 3},
		{"case 3", "0", 0},
		{"case 4", "06", 0},
		{"case 5", "10", 1},
		{"case 6", "101", 1},
		{"case 7", "100", 0},
		{"case 8", "1010", 1},
		{"case 9", "10100", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			actual := numDecodings(tc.input)
			if actual != tc.expected {
				t.Errorf(`testcase: %s
				input = %s 
				expected %v but got %v`, tc.caseName, tc.input, tc.expected, actual)
			}
		})
	}
}
