package potato_bugs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildList(names []string, cycleInd int) *PotatoNode {
	if len(names) == 0 {
		return nil
	}

	head := &PotatoNode{Name: names[0]}
	node := head
	nodes := []*PotatoNode{head}
	for ind := 1; ind < len(names); ind++ {
		node.Next = &PotatoNode{Name: names[ind]}
		node = node.Next
		nodes = append(nodes, node)
	}

	if cycleInd >= 0 && cycleInd < len(nodes) {
		node.Next = nodes[cycleInd]
	}
	return head
}

func TestHasPotatoCycle(t *testing.T) {
	type testCase struct {
		name           string
		inputNames     []string
		inputCycleInd  int
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "empty_list",
			inputNames:     nil,
			inputCycleInd:  -1,
			expectedOutput: false,
		},
		{
			name:           "one_element_no_cycle",
			inputNames:     []string{"Jack"},
			inputCycleInd:  -1,
			expectedOutput: false,
		},
		{
			name:           "one_element_cycle",
			inputNames:     []string{"Jack"},
			inputCycleInd:  0,
			expectedOutput: true,
		},
		{
			name:           "two_elements_no_cycle",
			inputNames:     []string{"Jack", "John"},
			inputCycleInd:  -1,
			expectedOutput: false,
		},
		{
			name:           "two_elements_cycle",
			inputNames:     []string{"Jack", "John"},
			inputCycleInd:  0,
			expectedOutput: true,
		},
		{
			name:           "many_elements_no_cycle",
			inputNames:     []string{"Jack", "John", "Maria", "Mia", "Amelia"},
			inputCycleInd:  -1,
			expectedOutput: false,
		},
		{
			name:           "many_elements_cycle",
			inputNames:     []string{"Jack", "John", "Maria", "Mia", "Amelia"},
			inputCycleInd:  2,
			expectedOutput: true,
		},
		{
			name:           "many_elements_last_element_cycle",
			inputNames:     []string{"Jack", "John", "Maria", "Mia", "Amelia"},
			inputCycleInd:  4,
			expectedOutput: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.inputNames, tc.inputCycleInd)
			result := hasPotatoCycle(head)
			assert.Equal(t, result, tc.expectedOutput)
		})
	}
}
