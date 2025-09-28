package tests

import (
	"Practice/utils"
	"testing"
)

func TestCreateCircularLinkedList(t *testing.T) {
	tests := []struct {
		name        string
		values      []int
		pos         int
		hasCycle    bool
		cycleLength int
		description string
	}{
		{
			name:        "Regular List (No Cycle)",
			values:      []int{1, 2, 3, 4, 5},
			pos:         -1,
			hasCycle:    false,
			cycleLength: 0,
			description: "pos = -1 means no cycle",
		},
		{
			name:        "Single Node No Cycle",
			values:      []int{1},
			pos:         -1,
			hasCycle:    false,
			cycleLength: 0,
			description: "Single node with no cycle",
		},
		{
			name:        "Single Node Self Cycle",
			values:      []int{1},
			pos:         0,
			hasCycle:    true,
			cycleLength: 1,
			description: "Single node pointing to itself",
		},
		{
			name:        "Two Nodes Cycle",
			values:      []int{1, 2},
			pos:         0,
			hasCycle:    true,
			cycleLength: 2,
			description: "Two nodes forming a cycle",
		},
		{
			name:        "Three Nodes Cycle at Position 1",
			values:      []int{1, 2, 3},
			pos:         1,
			hasCycle:    true,
			cycleLength: 2,
			description: "Three nodes with cycle at position 1",
		},
		{
			name:        "Four Nodes Cycle at Position 1",
			values:      []int{3, 2, 0, -4},
			pos:         1,
			hasCycle:    true,
			cycleLength: 3,
			description: "Four nodes with cycle at position 1",
		},
		{
			name:        "Large List No Cycle",
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			pos:         -1,
			hasCycle:    false,
			cycleLength: 0,
			description: "Large list with no cycle",
		},
		{
			name:        "Large List with Cycle",
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			pos:         3,
			hasCycle:    true,
			cycleLength: 6,
			description: "Large list with cycle at position 3",
		},
		{
			name:        "Empty List",
			values:      []int{},
			pos:         -1,
			hasCycle:    false,
			cycleLength: 0,
			description: "Empty list should return nil",
		},
		{
			name:        "Cycle at Last Position",
			values:      []int{1, 2, 3, 4, 5},
			pos:         4,
			hasCycle:    true,
			cycleLength: 1,
			description: "Cycle at the last position (tail points to itself)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create circular linked list
			head := utils.CreateCircularLinkedList(tt.values, tt.pos)

			// Test empty list case
			if len(tt.values) == 0 {
				if head != nil {
					t.Errorf("Expected nil for empty list, got %v", head)
				}
				return
			}

			// Test cycle detection
			hasCycle := utils.HasCycle(head)
			if hasCycle != tt.hasCycle {
				t.Errorf("HasCycle() = %v, expected %v", hasCycle, tt.hasCycle)
				t.Errorf("Test: %s", tt.description)
				t.Errorf("Input: values=%v, pos=%d", tt.values, tt.pos)
			}

			// Test cycle length (only if there's a cycle)
			if tt.hasCycle {
				cycleLength := utils.GetCycleLength(head)
				if cycleLength != tt.cycleLength {
					t.Errorf("GetCycleLength() = %d, expected %d", cycleLength, tt.cycleLength)
					t.Errorf("Test: %s", tt.description)
					t.Errorf("Input: values=%v, pos=%d", tt.values, tt.pos)
				}
			}
		})
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int
		expected bool
	}{
		{"No Cycle", []int{1, 2, 3, 4}, -1, false},
		{"Has Cycle", []int{1, 2, 3, 4}, 1, true},
		{"Empty List", []int{}, -1, false},
		{"Single Node No Cycle", []int{1}, -1, false},
		{"Single Node Self Cycle", []int{1}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.CreateCircularLinkedList(tt.values, tt.pos)
			result := utils.HasCycle(head)
			if result != tt.expected {
				t.Errorf("HasCycle() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestGetCycleLength(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int
		expected int
	}{
		{"No Cycle", []int{1, 2, 3, 4}, -1, 0},
		{"Cycle Length 2", []int{1, 2}, 0, 2},
		{"Cycle Length 3", []int{1, 2, 3}, 0, 3},
		{"Cycle Length 4", []int{1, 2, 3, 4}, 0, 4},
		{"Self Cycle", []int{1}, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.CreateCircularLinkedList(tt.values, tt.pos)
			result := utils.GetCycleLength(head)
			if result != tt.expected {
				t.Errorf("GetCycleLength() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
