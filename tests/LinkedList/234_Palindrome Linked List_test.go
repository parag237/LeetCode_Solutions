package tests

import (
	"Practice/code/LinkedList"
	"Practice/utils"
	"testing"
)

// TestIsPalindrome - Comprehensive test cases for palindrome detection
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expected    bool
		description string
	}{
		// Basic palindrome cases
		{
			name:        "Basic Palindrome - Even Length",
			input:       []int{1, 2, 2, 1},
			expected:    true,
			description: "Simple palindrome with even number of elements",
		},
		{
			name:        "Basic Palindrome - Odd Length",
			input:       []int{1, 2, 3, 2, 1},
			expected:    true,
			description: "Simple palindrome with odd number of elements",
		},

		// Non-palindrome cases
		{
			name:        "Non-Palindrome - Even Length",
			input:       []int{1, 2},
			expected:    false,
			description: "Simple non-palindrome with even number of elements",
		},
		{
			name:        "Non-Palindrome - Odd Length",
			input:       []int{1, 2, 3, 4, 5},
			expected:    false,
			description: "Simple non-palindrome with odd number of elements",
		},

		// Edge cases
		{
			name:        "Single Element",
			input:       []int{1},
			expected:    true,
			description: "Single element is considered a palindrome",
		},
		{
			name:        "Empty List",
			input:       []int{},
			expected:    true,
			description: "Empty list is considered a palindrome",
		},
		{
			name:        "Two Identical Elements",
			input:       []int{5, 5},
			expected:    true,
			description: "Two identical elements form a palindrome",
		},
		{
			name:        "Two Different Elements",
			input:       []int{3, 7},
			expected:    false,
			description: "Two different elements do not form a palindrome",
		},

		// Negative numbers
		{
			name:        "Palindrome with Negative Numbers",
			input:       []int{-1, 2, -1},
			expected:    true,
			description: "Palindrome containing negative numbers",
		},
		{
			name:        "Non-Palindrome with Negative Numbers",
			input:       []int{-1, 2, -3},
			expected:    false,
			description: "Non-palindrome containing negative numbers",
		},
		{
			name:        "All Negative Numbers - Palindrome",
			input:       []int{-5, -3, -5},
			expected:    true,
			description: "Palindrome with all negative numbers",
		},

		// Zero cases
		{
			name:        "Palindrome with Zeros",
			input:       []int{0, 1, 0},
			expected:    true,
			description: "Palindrome containing zeros",
		},
		{
			name:        "All Zeros",
			input:       []int{0, 0, 0, 0},
			expected:    true,
			description: "All zeros form a palindrome",
		},
		{
			name:        "Non-Palindrome with Zeros",
			input:       []int{0, 1, 2, 0},
			expected:    false,
			description: "Non-palindrome containing zeros",
		},

		// Large palindromes
		{
			name:        "Large Palindrome - 7 Elements",
			input:       []int{1, 2, 3, 4, 3, 2, 1},
			expected:    true,
			description: "Large palindrome with 7 elements",
		},
		{
			name:        "Large Non-Palindrome - 6 Elements",
			input:       []int{1, 2, 3, 4, 5, 6},
			expected:    false,
			description: "Large non-palindrome with 6 elements",
		},

		// Mixed cases
		{
			name:        "Mixed Positive and Negative - Palindrome",
			input:       []int{-2, 0, 2, 0, -2},
			expected:    true,
			description: "Palindrome with mixed positive, negative, and zero values",
		},
		{
			name:        "Mixed Positive and Negative - Non-Palindrome",
			input:       []int{-1, 0, 1, 2},
			expected:    false,
			description: "Non-palindrome with mixed positive, negative, and zero values",
		},

		// Special cases
		{
			name:        "Palindrome with Repeated Elements",
			input:       []int{1, 1, 1, 1, 1},
			expected:    true,
			description: "Palindrome with all identical elements",
		},
		{
			name:        "Alternating Pattern - Palindrome",
			input:       []int{1, 2, 1, 2, 1},
			expected:    true,
			description: "Alternating pattern that forms a palindrome",
		},
		{
			name:        "Alternating Pattern - Non-Palindrome",
			input:       []int{1, 2, 1, 2},
			expected:    false,
			description: "Alternating pattern that does not form a palindrome",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the linked list from input
			head := utils.CreateLinkedList(tt.input)

			// Call the isPalindrome function
			// Note: This assumes isPalindrome is accessible from main package
			// You may need to move this function to a shared package
			result := LinkedList.IsPalindrome(head)

			// Check result
			if result != tt.expected {
				t.Errorf("isPalindrome() = %v, expected %v", result, tt.expected)
				t.Errorf("Test: %s", tt.description)
				t.Errorf("Input: %v", tt.input)
			}
		})
	}
}
