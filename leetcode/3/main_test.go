package main

import (
	"testing"
)

func TestValidations(t *testing.T) {
	type TestCases struct {
		input    string
		expected int
	}

	testCases := []TestCases{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    " ",
			expected: 1,
		},
		{
			input:    "tmmzuxt",
			expected: 5,
		},
		{
			input:    "dvdf",
			expected: 3,
		},
	}

	for _, testCase := range testCases {
		expected := testCase.expected
		input := testCase.input

		r := lengthOfLongestSubstring(input)
		if r != expected {
			t.Errorf("Input ['%s']. Expected %d, Received %d\n", input, expected, r)
		}
	}
}
