package walnut_iq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckWalnutGenius(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedOutput []int
	}
	testCases := []testCase{
		{
			name:           "four_elements",
			input:          []int{14, 75, 24, 86},
			expectedOutput: []int{61, 41, 58, 37},
		},
		{
			name:           "three_elements",
			input:          []int{15, 23, 48},
			expectedOutput: []int{35, 31, 19},
		},
		{
			name:           "two_elements",
			input:          []int{15, 23},
			expectedOutput: []int{23, 15},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkWalnutGenius(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
