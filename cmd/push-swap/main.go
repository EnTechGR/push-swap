package main

import (
	"fmt"
	"os"
	"push-swap/internal/parser"
	"push-swap/internal/solver"
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
	
	// Create solver and solve
	s := solver.NewSolver(numbers)
	operations := s.Solve()
	
	// Output operations
	for _, op := range operations {
		fmt.Println(op)
	}
}
