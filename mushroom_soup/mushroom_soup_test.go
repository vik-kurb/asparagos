package mushroom_soup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	type testCase struct {
		name           string
		inputMushrooms []int
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "two_identical_mushrooms",
			inputMushrooms: []int{6, 6},
			expectedOutput: true,
		},
		{
			name:           "different_mushrooms_count",
			inputMushrooms: []int{10, 2, 3, 5},
			expectedOutput: true,
		},
		{
			name:           "odd_mushrooms_count",
			inputMushrooms: []int{7, 2, 4, 3, 2},
			expectedOutput: true,
		},
		{
			name:           "even_mushrooms_count_impossible",
			inputMushrooms: []int{6, 8},
			expectedOutput: false,
		},
		{
			name:           "odd_mushrooms_count_impossible",
			inputMushrooms: []int{6, 7, 9},
			expectedOutput: false,
		},
		{
			name:           "odd_mushrooms_sum",
			inputMushrooms: []int{6, 7},
			expectedOutput: false,
		},
		{
			name:           "single_mushroom",
			inputMushrooms: []int{9},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canPartition(tc.inputMushrooms)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
