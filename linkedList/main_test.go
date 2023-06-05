package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LinkedListTestSuite struct {
	suite.Suite
}

func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}

func (s *LinkedListTestSuite) TestAdd() {
	type TestCases struct {
		input               []int
		expectedLen         int
		expectedListAsArray []int
	}

	testCases := []TestCases{
		{
			input:               []int{1},
			expectedLen:         1,
			expectedListAsArray: []int{1},
		},
		{
			input:               []int{1, 9, 7},
			expectedLen:         3,
			expectedListAsArray: []int{1, 9, 7},
		},
		{
			input:               []int{},
			expectedLen:         0,
			expectedListAsArray: nil,
		},
	}
	for _, testCase := range testCases {
		list := NewLinkedList[int]()
		for _, input := range testCase.input {
			list.Add(input)
		}
		assert.Equal(s.T(), list.Len(), testCase.expectedLen)
		assert.Equal(s.T(), list.AsArray(), testCase.expectedListAsArray)
	}
}
