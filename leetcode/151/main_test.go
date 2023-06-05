package leetcode

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
		s        string
		expected string
	}

	testCases := []TestCases{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for _, testCase := range testCases {
		r := reverseWords(testCase.s)
		s.Assert().Equal(testCase.expected, r)
	}
}
