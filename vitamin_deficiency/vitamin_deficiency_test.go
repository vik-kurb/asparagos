package vitamin_deficiency

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindVitaminsSum(t *testing.T) {
	type testCase struct {
		name            string
		inputVitamins   []int
		inputTargetSum  int
		expectedOutputs [][]int
	}
	testCases := []testCase{
		{
			name:            "found_pair",
			inputVitamins:   []int{26, 73, 14, 61, 100},
			inputTargetSum:  134,
			expectedOutputs: [][]int{{1, 3}},
		},
		{
			name:            "found_one_of_pairs",
			inputVitamins:   []int{26, 73, 14, 61, 99},
			inputTargetSum:  99,
			expectedOutputs: [][]int{{0, 1}, {2, 4}},
		},
		{
			name:            "found_pair_similar_elements",
			inputVitamins:   []int{26, 73, 26, 61, 100},
			inputTargetSum:  52,
			expectedOutputs: [][]int{{0, 2}},
		},
		{
			name:            "no_pair",
			inputVitamins:   []int{26, 73, 14, 61, 100},
			inputTargetSum:  10,
			expectedOutputs: nil,
		},
		{
			name:            "no_pair_no_duplicate_elements",
			inputVitamins:   []int{26, 73, 14, 61, 100},
			inputTargetSum:  52,
			expectedOutputs: nil,
		},
		{
			name:            "empty_input",
			inputVitamins:   []int{},
			inputTargetSum:  52,
			expectedOutputs: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findVitaminsSum(tc.inputVitamins, tc.inputTargetSum)
			if tc.expectedOutputs == nil {
				assert.Equal(t, []int(nil), result)
			} else {
				foundExpectedOutput := false
				for _, expectedOutput := range tc.expectedOutputs {
					if slices.Equal(expectedOutput, result) {
						foundExpectedOutput = true
					}
				}
				assert.True(t, foundExpectedOutput)
			}
		})
	}
}
