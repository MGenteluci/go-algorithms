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
			s:        "hello",
			expected: "holle",
		},
		{
			s:        "leetcode",
			expected: "leotcede",
		},
	}

	for _, testCase := range testCases {
		r := reverseVowels(testCase.s)
		s.Assert().Equal(testCase.expected, r)
	}
}
