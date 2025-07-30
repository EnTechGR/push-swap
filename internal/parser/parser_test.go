package parser

import (
	"reflect"
	"testing"
)

func TestParseArguments(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []int
		hasError bool
	}{
		{
			name:     "Empty arguments",
			args:     []string{},
			expected: []int{},
			hasError: false,
		},
		{
			name:     "Single number",
			args:     []string{"42"},
			expected: []int{42},
			hasError: false,
		},
		{
			name:     "Multiple arguments",
			args:     []string{"1", "2", "3"},
			expected: []int{1, 2, 3},
			hasError: false,
		},
		{
			name:     "Space-separated string",
			args:     []string{"2 1 3 6 5 8"},
			expected: []int{2, 1, 3, 6, 5, 8},
			hasError: false,
		},
		{
			name:     "Mixed format",
			args:     []string{"1 2", "3", "4 5"},
			expected: []int{1, 2, 3, 4, 5},
			hasError: false,
		},
		{
			name:     "Negative numbers",
			args:     []string{"-1", "0", "1"},
			expected: []int{-1, 0, 1},
			hasError: false,
		},
		{
			name:     "Invalid integer",
			args:     []string{"1", "not_a_number", "3"},
			expected: nil,
			hasError: true,
		},
		{
			name:     "Duplicate numbers",
			args:     []string{"1", "2", "1"},
			expected: nil,
			hasError: true,
		},
		{
			name:     "Large numbers",
			args:     []string{"2147483647", "-2147483648"},
			expected: []int{2147483647, -2147483648},
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseArguments(tt.args)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCheckDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		hasError bool
	}{
		{
			name:     "No duplicates",
			numbers:  []int{1, 2, 3, 4, 5},
			hasError: false,
		},
		{
			name:     "Empty slice",
			numbers:  []int{},
			hasError: false,
		},
		{
			name:     "Single element",
			numbers:  []int{42},
			hasError: false,
		},
		{
			name:     "Duplicates present",
			numbers:  []int{1, 2, 3, 2, 4},
			hasError: true,
		},
		{
			name:     "All same elements",
			numbers:  []int{5, 5, 5},
			hasError: true,
		},
		{
			name:     "Negative duplicates",
			numbers:  []int{-1, -2, -1},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDuplicates(tt.numbers)
			
			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestParseOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
		hasError bool
	}{
		{
			name:     "Valid operations",
			input:    []string{"sa", "pb", "rra"},
			expected: []string{"sa", "pb", "rra"},
			hasError: false,
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: []string{},
			hasError: false,
		},
		{
			name:     "Operations with whitespace",
			input:    []string{" sa ", " pb\n", "rra "},
			expected: []string{"sa", "pb", "rra"},
			hasError: false,
		},
		{
			name:     "Empty strings in input",
			input:    []string{"sa", "", "pb", "   ", "rra"},
			expected: []string{"sa", "pb", "rra"},
			hasError: false,
		},
		{
			name:     "Invalid operation",
			input:    []string{"sa", "invalid_op", "pb"},
			expected: nil,
			hasError: true,
		},
		{
			name:     "All valid operations",
			input:    []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"},
			expected: []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"},
			hasError: false,
		},
		{
			name:     "Case sensitive",
			input:    []string{"SA", "pb"},
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseOperations(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}