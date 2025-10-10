package lemon_squeezy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBestTeamIQ(t *testing.T) {
	type testCase struct {
		name           string
		inputIQs       []int
		inputAges      []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "all_lemons",
			inputIQs:       []int{1, 2, 3},
			inputAges:      []int{10, 15, 20},
			expectedOutput: 6,
		},
		{
			name:           "one_element",
			inputIQs:       []int{7},
			inputAges:      []int{26},
			expectedOutput: 7,
		},
		{
			name:           "two_elements_all_in_team",
			inputIQs:       []int{16, 32},
			inputAges:      []int{5, 23},
			expectedOutput: 48,
		},
		{
			name:           "two_elements_one_in_team",
			inputIQs:       []int{16, 32},
			inputAges:      []int{46, 23},
			expectedOutput: 32,
		},
		{
			name:           "many_elements_one_lemon_in_team",
			inputIQs:       []int{97, 10, 4, 12, 5, 20},
			inputAges:      []int{1, 2, 3, 4, 5, 6},
			expectedOutput: 97,
		},
		{
			name:           "not_all_lemons_in_team",
			inputIQs:       []int{17, 10, 4, 12, 5, 20},
			inputAges:      []int{1, 2, 3, 4, 5, 6},
			expectedOutput: 42,
		},
		{
			name:           "equal_ages",
			inputIQs:       []int{10, 20, 30},
			inputAges:      []int{5, 5, 5},
			expectedOutput: 60,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findBestTeamIQ(tc.inputIQs, tc.inputAges)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
