package main

import (
	"fmt"
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
		numbers  []int
		target   int
		expected []int
	}

	testCases := []TestCases{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	for _, testCase := range testCases {
		r := twoSum(testCase.numbers, testCase.target)
		fmt.Printf("r: %v\n", r)
		s.ElementsMatch(r, testCase.expected)
	}
}
