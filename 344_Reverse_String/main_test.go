package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	var testCasesStruct = []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
		{[]byte("12345"), []byte("54321")},
		{[]byte("1"), []byte("1")},
		{[]byte(""), []byte("")},
	}

	for _, testCase := range testCasesStruct {
		reverseString(testCase.input)
		input := string(testCase.input)
		expected := string(testCase.expected)
		if input != expected {
			t.Errorf(`Expected: %s 
			but got: %s`, testCase.expected, testCase.input)
		}
	}
}
