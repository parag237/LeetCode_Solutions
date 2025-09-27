package tests

import (
	"Practice/code/LinkedList"
	"Practice/utils"
	"testing"
)

func TestReorderList(t *testing.T) {
	testCases := []struct {
		name        string
		input       []int
		output      []int
		description string
	}{
		{
			name:        "Even number of nodes",
			input:       []int{1, 2, 3, 4},
			output:      []int{1, 4, 2, 3},
			description: "This is the Ideal case",
		},
		{
			name:   "Odd number of nodes",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 5, 2, 4, 3},
		},
		{
			name:   "empty input",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "single input",
			input:  []int{1},
			output: []int{1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.CreateLinkedList(tt.input)
			expectedOutput := utils.CreateLinkedList(tt.output)

			LinkedList.ReorderList(head)
			if !utils.IsDuplicateList(head, expectedOutput) {
				t.Errorf("Invalid Output, Expected: %v", tt.output)
			}
		})
	}
}
