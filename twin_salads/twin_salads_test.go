package twin_salads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreSaladsTwin(t *testing.T) {
	type testCase struct {
		name           string
		input          []string
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "twin_even_length",
			input:          []string{"cucumber", "tomato", "pepper", "pepper", "tomato", "cucumber"},
			expectedOutput: true,
		},
		{
			name:           "twin_odd_length",
			input:          []string{"cucumber", "tomato", "pepper", "tomato", "cucumber"},
			expectedOutput: true,
		},
		{
			name:           "twin_one_element",
			input:          []string{"cucumber"},
			expectedOutput: true,
		},
		{
			name:           "twin_two_elements",
			input:          []string{"cucumber", "cucumber"},
			expectedOutput: true,
		},
		{
			name:           "twin_empty",
			input:          []string{},
			expectedOutput: true,
		},
		{
			name:           "twin_ignore_case",
			input:          []string{"cUcuMber", "TOMATO", "pepper", "peppeR", "tomato", "cucumber"},
			expectedOutput: true,
		},
		{
			name:           "not_twin_even_length",
			input:          []string{"cucumber", "tomato", "pepper", "lettuce"},
			expectedOutput: false,
		},
		{
			name:           "not_twin_odd_length",
			input:          []string{"cucumber", "tomato", "lettuce"},
			expectedOutput: false,
		},
		{
			name:           "not_twin_two_elements",
			input:          []string{"cucumber", "tomato"},
			expectedOutput: false,
		},
		{
			name:           "not_twin_incorrect_order",
			input:          []string{"cucumber", "tomato", "cucumber", "tomato"},
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := areSaladsTwin(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
