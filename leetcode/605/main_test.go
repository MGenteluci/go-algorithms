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
		flowerbed []int
		n         int
		expected  bool
	}

	testCases := []TestCases{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
			expected:  true,
		},
	}

	for _, testCase := range testCases {
		r := canPlaceFlowers(testCase.flowerbed, testCase.n)
		s.Assert().Equal(testCase.expected, r)
	}
}
