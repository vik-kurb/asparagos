package apple_escape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSaveAllApples(t *testing.T) {
	type testCase struct {
		name           string
		inputCapacity  int
		inputTrips     []Trip
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "trips_one_by_one",
			inputCapacity:  7,
			inputTrips:     []Trip{{from: 0, to: 2, num: 7}, {from: 2, to: 3, num: 7}, {from: 3, to: 7, num: 7}},
			expectedOutput: true,
		},
		{
			name:           "has_intersections",
			inputCapacity:  5,
			inputTrips:     []Trip{{from: 2, to: 7, num: 1}, {from: 0, to: 2, num: 3}, {from: 1, to: 3, num: 2}},
			expectedOutput: true,
		},
		{
			name:           "no_intersections",
			inputCapacity:  10,
			inputTrips:     []Trip{{from: 4, to: 8, num: 2}, {from: 0, to: 2, num: 5}, {from: 10, to: 15, num: 1}},
			expectedOutput: true,
		},
		{
			name:           "multiple_intersections",
			inputCapacity:  5,
			inputTrips:     []Trip{{from: 1, to: 4, num: 2}, {from: 2, to: 3, num: 2}, {from: 0, to: 5, num: 1}},
			expectedOutput: true,
		},
		{
			name:           "one_trip",
			inputCapacity:  2,
			inputTrips:     []Trip{{from: 0, to: 2, num: 2}},
			expectedOutput: true,
		},
		{
			name:           "one_trip_impossible",
			inputCapacity:  2,
			inputTrips:     []Trip{{from: 0, to: 2, num: 6}},
			expectedOutput: false,
		},
		{
			name:           "no_intersections_impossible",
			inputCapacity:  2,
			inputTrips:     []Trip{{from: 4, to: 8, num: 2}, {from: 0, to: 2, num: 5}},
			expectedOutput: false,
		},
		{
			name:           "multiple_intersections_impossible",
			inputCapacity:  4,
			inputTrips:     []Trip{{from: 1, to: 4, num: 2}, {from: 2, to: 3, num: 2}, {from: 0, to: 5, num: 1}},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canSaveAllApples(tc.inputTrips, tc.inputCapacity)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
