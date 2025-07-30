package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseArguments parses command line arguments into integers
func ParseArguments(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}
	
	var numbers []int
	var allArgs []string
	
	// Handle the case where all numbers are in a single string
	for _, arg := range args {
		// Split by spaces to handle "2 1 3 6 5 8" format
		parts := strings.Fields(arg)
		allArgs = append(allArgs, parts...)
	}
	
	// Convert strings to integers
	for _, str := range allArgs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", str)
		}
		numbers = append(numbers, num)
	}
	
	// Check for duplicates
	if err := checkDuplicates(numbers); err != nil {
		return nil, err
	}
	
	return numbers, nil
}

// checkDuplicates checks if there are duplicate numbers in the slice
func checkDuplicates(numbers []int) error {
	seen := make(map[int]bool)
	for _, num := range numbers {
		if seen[num] {
			return fmt.Errorf("duplicate number found: %d", num)
		}
		seen[num] = true
	}
	return nil
}

// ParseOperations parses operation strings into Operation types
func ParseOperations(operationStrings []string) ([]string, error) {
	validOps := map[string]bool{
		"sa": true, "sb": true, "ss": true,
		"pa": true, "pb": true,
		"ra": true, "rb": true, "rr": true,
		"rra": true, "rrb": true, "rrr": true,
	}
	
	var operations []string
	for _, opStr := range operationStrings {
		opStr = strings.TrimSpace(opStr)
		if opStr == "" {
			continue
		}
		if !validOps[opStr] {
			return nil, fmt.Errorf("invalid operation: %s", opStr)
		}
		operations = append(operations, opStr)
	}
	
	return operations, nil
}