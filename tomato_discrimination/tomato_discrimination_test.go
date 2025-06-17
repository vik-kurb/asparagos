package tomato_discrimination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTomato(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "tomato",
			input:          "tomato",
			expectedOutput: true,
		},
		{
			name:           "pomodoro",
			input:          "pomodoro",
			expectedOutput: true,
		},
		{
			name:           "not_tomato",
			input:          "cucumber",
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isTomato(tc.input)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}

func TestDiscriminateTomatoes(t *testing.T) {
	type testCase struct {
		name           string
		input          []string
		expectedOutput []string
	}
	testCases := []testCase{
		{
			name:           "no_tomato",
			input:          []string{"cucumber", "pepper", "lettuce"},
			expectedOutput: []string{"cucumber", "pepper", "lettuce"},
		},
		{
			name:           "with_tomato",
			input:          []string{"cucumber", "tomato", "pepper", "lettuce", "tomato"},
			expectedOutput: []string{"cucumber", "pepper", "lettuce", "tomato", "tomato"},
		},
		{
			name:           "with_pomodoro",
			input:          []string{"cucumber", "pepper", "pomodoro", "pomodoro", "lettuce"},
			expectedOutput: []string{"cucumber", "pepper", "lettuce", "tomato", "tomato"},
		},
		{
			name:           "with_tomato_and_pomodoro",
			input:          []string{"cucumber", "tomato", "pomodoro", "pepper", "pomodoro", "lettuce", "tomato"},
			expectedOutput: []string{"cucumber", "pepper", "lettuce", "tomato", "tomato", "tomato", "tomato"},
		},
		{
			name:           "only_tomato_and_pomodoro",
			input:          []string{"tomato", "pomodoro", "pomodoro", "tomato", "tomato"},
			expectedOutput: []string{"tomato", "tomato", "tomato", "tomato", "tomato"},
		},
		{
			name:           "empty",
			input:          []string{},
			expectedOutput: []string{},
		},
		{
			name:           "one_element_no_tomato",
			input:          []string{"pepper"},
			expectedOutput: []string{"pepper"},
		},
		{
			name:           "one_element_tomato",
			input:          []string{"tomato"},
			expectedOutput: []string{"tomato"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vegetables := make([]string, len(tc.input))
			copy(vegetables, tc.input)
			discriminateTomatoes(vegetables)
			assert.Equal(t, tc.expectedOutput, vegetables)
		})
	}
}
