package strawberry_fever

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanEatAllStrawberries(t *testing.T) {
	type testCase struct {
		name           string
		storeStock     []int
		wayCost        []int
		expectedOutput int
	}
	testCases := []testCase{
		{
			name:           "possible",
			storeStock:     []int{1, 2, 3, 4, 5},
			wayCost:        []int{3, 4, 5, 1, 2},
			expectedOutput: 3,
		},
		{
			name:           "impossible",
			storeStock:     []int{2, 3, 4},
			wayCost:        []int{3, 4, 3},
			expectedOutput: -1,
		},
		{
			name:           "start_at_the_end",
			storeStock:     []int{2, 1, 1, 10},
			wayCost:        []int{2, 2, 2, 2},
			expectedOutput: 3,
		},
		{
			name:           "start_at_the_beginning",
			storeStock:     []int{5, 1, 1, 1},
			wayCost:        []int{1, 2, 2, 2},
			expectedOutput: 0,
		},
		{
			name:           "start_at_the_beginning",
			storeStock:     []int{6, 0, 6, 0, 6},
			wayCost:        []int{4, 4, 4, 4, 4},
			expectedOutput: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canEatAllStrawberries(tc.storeStock, tc.wayCost)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
