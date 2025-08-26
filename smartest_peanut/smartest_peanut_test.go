package peanuts_ego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSmartestPeanut(t *testing.T) {
	type testCase struct {
		name           string
		inputPeanuts   []int
		inputK         int
		expectedOutput []int
	}
	testCases := []testCase{
		{
			name:           "large_window",
			inputPeanuts:   []int{1, 6, 3, 4, 1, 2, 9},
			inputK:         5,
			expectedOutput: []int{6, 6, 9},
		},
		{
			name:           "small_window",
			inputPeanuts:   []int{1, 6, 3, 4, 1, 2, 9},
			inputK:         2,
			expectedOutput: []int{6, 6, 4, 4, 2, 9},
		},
		{
			name:           "one_element_window",
			inputPeanuts:   []int{20, 15, 13, 9, 6, 2},
			inputK:         1,
			expectedOutput: []int{20, 15, 13, 9, 6, 2},
		},
		{
			name:           "one_element",
			inputPeanuts:   []int{4},
			inputK:         1,
			expectedOutput: []int{4},
		},
		{
			name:           "one_window",
			inputPeanuts:   []int{4, 10, 5, 2, 65, 2},
			inputK:         6,
			expectedOutput: []int{65},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findSmartestPeanut(tc.inputPeanuts, tc.inputK)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
