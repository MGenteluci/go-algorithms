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
		str1     string
		str2     string
		expected string
	}

	testCases := []TestCases{
		{
			str1:     "ABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			str1:     "ABABAB",
			str2:     "ABAB",
			expected: "AB",
		},
		{
			str1:     "LEET",
			str2:     "CODE",
			expected: "",
		},
	}

	for _, testCase := range testCases {
		r := gcdOfStrings(testCase.str1, testCase.str2)
		s.Assert().Equal(testCase.expected, r)
	}
}
