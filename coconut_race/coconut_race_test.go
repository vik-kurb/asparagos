package coconut_race

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCoconutBouquets(t *testing.T) {
	type testCase struct {
		name           string
		inputTarget    int
		inputPosition  []int
		inputSpeed     []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "all_in_one_bouquet",
			inputTarget:    10,
			inputPosition:  []int{2, 0, 4},
			inputSpeed:     []int{3, 7, 1},
			expectedOutput: 1,
		},
		{
			name:           "all_finish_alone",
			inputTarget:    12,
			inputPosition:  []int{2, 7, 0, 3},
			inputSpeed:     []int{2, 5, 1, 3},
			expectedOutput: 4,
		},
		{
			name:           "several_bouquets",
			inputTarget:    15,
			inputPosition:  []int{0, 1, 3, 7, 9, 12},
			inputSpeed:     []int{2, 1, 5, 20, 10, 2},
			expectedOutput: 3,
		},
		{
			name:           "merge_at_finish",
			inputTarget:    6,
			inputPosition:  []int{0, 3, 2},
			inputSpeed:     []int{6, 3, 4},
			expectedOutput: 1,
		},
		{
			name:           "empty",
			inputTarget:    6,
			inputPosition:  []int{},
			inputSpeed:     []int{},
			expectedOutput: 0,
		},
		{
			name:           "same_position",
			inputTarget:    6,
			inputPosition:  []int{1, 1},
			inputSpeed:     []int{3, 5},
			expectedOutput: 1,
		},
		{
			name:           "one_coconut",
			inputTarget:    6,
			inputPosition:  []int{3},
			inputSpeed:     []int{7},
			expectedOutput: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countCoconutBouquets(tc.inputTarget, tc.inputPosition, tc.inputSpeed)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
