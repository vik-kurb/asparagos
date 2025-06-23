package fruit_sabotage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFruitKing(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "two_variants",
			input:          []int{1, 2, 1, 2, 1},
			expectedOutput: 1,
		},
		{
			name:           "many_variants",
			input:          []int{1, 2, 1, 1, 2, 4, 1, 1, 8},
			expectedOutput: 1,
		},
		{
			name:           "one_variant",
			input:          []int{15, 15, 15, 15, 15, 15},
			expectedOutput: 15,
		},
		{
			name:           "one_element",
			input:          []int{42},
			expectedOutput: 42,
		},
		{
			name:           "result_at_the_beginning",
			input:          []int{67, 67, 67, 67, 67, 67, 52, 52, 31},
			expectedOutput: 67,
		},
		{
			name:           "result_at_the_end",
			input:          []int{25, 65, 100, 7, 7, 7, 7, 7},
			expectedOutput: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findFruitKing(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
