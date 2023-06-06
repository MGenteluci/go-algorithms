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
		nums     []int
		expected []int
	}

	testCases := []TestCases{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for _, testCase := range testCases {
		r := productExceptSelf(testCase.nums)
		s.Assert().Equal(testCase.expected, r)
	}
}
