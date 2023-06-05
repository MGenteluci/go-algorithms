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
		candies      []int
		extraCandies int
		expected     []bool
	}

	testCases := []TestCases{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			expected:     []bool{true, false, true},
		},
	}

	for _, testCase := range testCases {
		r := kidsWithCandies(testCase.candies, testCase.extraCandies)
		s.Assert().Equal(testCase.expected, r)
	}
}
