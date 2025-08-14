package worlds_biggest_gazpacho

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectTomatoes(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "take_odd_beds",
			input:          []int{2, 1, 4, 3, 6, 5, 8, 7},
			expectedOutput: 20,
		},
		{
			name:           "take_even_beds",
			input:          []int{2, 10, 4, 45, 6, 27, 8, 18},
			expectedOutput: 100,
		},
		{
			name:           "take_one_bed",
			input:          []int{3, 7, 2},
			expectedOutput: 7,
		},
		{
			name:           "first_and_last_beds",
			input:          []int{6, 3, 5, 10},
			expectedOutput: 16,
		},
		{
			name:           "one_element",
			input:          []int{15},
			expectedOutput: 15,
		},
		{
			name:           "two_elements",
			input:          []int{15, 20},
			expectedOutput: 20,
		},
		{
			name:           "zero_elements",
			input:          []int{0, 0, 0, 0, 0},
			expectedOutput: 0,
		},
		{
			name:           "empty",
			input:          []int{},
			expectedOutput: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := collectTomatoes(tc.input)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
