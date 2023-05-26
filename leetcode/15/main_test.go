package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) TestValidations() {
	type TestCases struct {
		input    []int
		expected [][]int
	}

	testCases := []TestCases{
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, testCase := range testCases {
		input := testCase.input

		r := threeSum(input)
		s.ElementsMatch(r, testCase.expected)
	}
}
