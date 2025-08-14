package salsa_festival

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAllVeggiesCome(t *testing.T) {
	type testCase struct {
		name              string
		inputVegNum       int
		inputRequirements [][]int
		expectedOutput    bool
	}
	testCases := []testCase{
		{
			name:              "two_vegetables_possible",
			inputVegNum:       2,
			inputRequirements: [][]int{{1, 0}},
			expectedOutput:    true,
		},
		{
			name:              "two_vegetables_impossible",
			inputVegNum:       2,
			inputRequirements: [][]int{{1, 0}, {0, 1}},
			expectedOutput:    false,
		},
		{
			name:              "empty",
			inputVegNum:       0,
			inputRequirements: [][]int{},
			expectedOutput:    true,
		},
		{
			name:              "one_vegetable",
			inputVegNum:       1,
			inputRequirements: [][]int{},
			expectedOutput:    true,
		},
		{
			name:              "nobody_can_come",
			inputVegNum:       5,
			inputRequirements: [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {0, 4}},
			expectedOutput:    false,
		},
		{
			name:              "not_all_can_come",
			inputVegNum:       8,
			inputRequirements: [][]int{{0, 1}, {0, 4}, {1, 2}, {1, 3}, {4, 3}, {6, 7}, {7, 6}},
			expectedOutput:    false,
		},
		{
			name:              "all_can_come",
			inputVegNum:       8,
			inputRequirements: [][]int{{0, 1}, {0, 4}, {1, 2}, {1, 3}, {4, 3}, {6, 7}},
			expectedOutput:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canAllVeggiesCome(tc.inputVegNum, tc.inputRequirements)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
