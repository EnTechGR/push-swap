package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"push-swap/internal/operations"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func main() {
	// Handle no arguments case
	if len(os.Args) < 2 {
		return
	}
	
	// Parse command line arguments
	numbers, err := parser.ParseArguments(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	
	// Handle empty input
	if len(numbers) == 0 {
		return
	}
	
	// Create stacks
	stackA := stack.NewStack(numbers)
	stackB := stack.NewEmptyStack()
	
	// Read operations from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var operationStrings []string
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			operationStrings = append(operationStrings, line)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	
	// Parse operations
	parsedOps, err := parser.ParseOperations(operationStrings)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	
	// Execute operations
	for _, opStr := range parsedOps {
		op := operations.Operation(opStr)
		if err := operations.ExecuteOperation(stackA, stackB, op); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	
	// Check if stack A is sorted and stack B is empty
	if stackA.IsSorted() && stackB.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}