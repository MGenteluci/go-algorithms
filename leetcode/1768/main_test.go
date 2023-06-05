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
		word1    string
		word2    string
		expected string
	}

	testCases := []TestCases{
		{
			word1:    "abc",
			word2:    "pqr",
			expected: "apbqcr",
		},
		{
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
		{
			word1:    "abcd",
			word2:    "pq",
			expected: "apbqcd",
		},
	}

	for _, testCase := range testCases {
		r := mergeAlternately(testCase.word1, testCase.word2)
		s.Assert().Equal(testCase.expected, r)
	}
}
