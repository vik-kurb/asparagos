package orange_seeds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSowTheSeedsOfDiscord(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "empty",
			input:          []int{},
			expectedOutput: 0,
		},
		{
			name:           "one_element",
			input:          []int{3},
			expectedOutput: 0,
		},
		{
			name:           "many_elements",
			input:          []int{3, 1, 5, 1, 1},
			expectedOutput: 2,
		},
		{
			name:           "many_elements_2",
			input:          []int{1, 3, 16, 1, 1, 5, 1},
			expectedOutput: 3,
		},
		{
			name:           "many_elements_3",
			input:          []int{1, 1, 1, 1, 1, 1, 1},
			expectedOutput: 6,
		},
		{
			name:           "many_elements_4",
			input:          []int{10, 1, 1, 1, 1, 1, 1},
			expectedOutput: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sowTheSeedsOfDiscord(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
