package main

import (
	"reflect"
	"testing"
)

// this problem need [x, y]
// X is the duplicated number in array
// Y is the mismatch number in array

func TestFindErrorNums(t *testing.T) {
	testCaseStruct := []struct {
		caseName string
		input    []int
		expected []int
	}{
		{
			caseName: "case1 - normal case 1",
			input:    []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		{
			caseName: "case2 - normal case 2",
			input:    []int{1, 1},
			expected: []int{1, 2},
		},
		{
			caseName: "case3 - normal case 3",
			input:    []int{2, 2},
			expected: []int{2, 1},
		},
		{
			caseName: "case4 - this case make me need sort!",
			input:    []int{3, 3, 1},
			expected: []int{3, 2},
		},
		{
			caseName: "case5 - i try to add break and remove it bcuz this case",
			input:    []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9},
			expected: []int{2, 10},
		},
	}

	for _, testCase := range testCaseStruct {
		actual := findErrorNums(testCase.input)
		if actual == nil || testCase.expected == nil {
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf(`
				caseName: %s, 
				actual: %v, 
				expected: %v`, testCase.caseName, actual, testCase.expected)
			}
			continue
		}

		// Check if slices have the same length
		if len(actual) != len(testCase.expected) {
			t.Errorf(`
				caseName: %s, 
				actual: %v, 
				expected: %v`, testCase.caseName, actual, testCase.expected)
			continue
		}

		// Iterate through the slices and compare elements
		for i := 0; i < len(actual); i++ {
			if actual[i] != testCase.expected[i] {
				t.Errorf(`
				caseName: %s, 
				actual: %v, 
				expected: %v`, testCase.caseName, actual, testCase.expected)
				continue
			}
		}
	}

}
