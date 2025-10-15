package watermelon_football

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWaysToMakeTeams(t *testing.T) {
	type testCase struct {
		name           string
		inputRating    []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "even_number_of_watermelons",
			inputRating:    []int{1, 3, 2, 1, 2, 2},
			expectedOutput: 2,
		},
		{
			name:           "odd_number_of_watermelons",
			inputRating:    []int{5, 7, 3, 8, 1},
			expectedOutput: 1,
		},
		{
			name:           "even_number_of_watermelons_zero_ways",
			inputRating:    []int{1, 2, 3, 4},
			expectedOutput: 0,
		},
		{
			name:           "odd_number_of_watermelons_zero_ways",
			inputRating:    []int{10, 6, 1},
			expectedOutput: 0,
		},
		{
			name:           "each_can_be_removed",
			inputRating:    []int{3, 3, 3, 3, 3},
			expectedOutput: 5,
		},
		{
			name:           "zero_ratings",
			inputRating:    []int{0, 1, 1, 0, 0, 0},
			expectedOutput: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countWaysToMakeTeams(tc.inputRating)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
