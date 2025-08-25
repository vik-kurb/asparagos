package sunflower_speaking_club

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSpeakingPartner(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedOutput []int
	}
	testCases := []testCase{
		{
			name:           "close_partners",
			input:          []int{1, 5, 7, 3, 4},
			expectedOutput: []int{1, 1, 0, 1, 0},
		},
		{
			name:           "distant_partners",
			input:          []int{10, 9, 8, 7, 11, 3},
			expectedOutput: []int{4, 3, 2, 1, 0, 0},
		},
		{
			name:           "no_partner",
			input:          []int{20, 15, 13, 9, 6, 2},
			expectedOutput: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name:           "empty",
			input:          []int{},
			expectedOutput: []int{},
		},
		{
			name:           "one_element",
			input:          []int{1},
			expectedOutput: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findSpeakingPartner(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
