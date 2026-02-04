package solver

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
	"testing"
)

func TestNewSolver(t *testing.T) {
	input := []int{3, 2, 1}
	solver := NewSolver(input)
	
	if solver == nil {
		t.Fatal("NewSolver should not return nil")
	}
	
	if solver.stackA.Size() != 3 {
		t.Errorf("Expected stack A size 3, got %d", solver.stackA.Size())
	}
	
	if !solver.stackB.IsEmpty() {
		t.Error("Stack B should be empty initially")
	}
	
	if len(solver.operations) != 0 {
		t.Error("Operations should be empty initially")
	}
}

func TestSolveAlreadySorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for already sorted stack, got %d", len(ops))
	}
}

func TestSolveEmptyStack(t *testing.T) {
	input := []int{}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for empty stack, got %d", len(ops))
	}
}

func TestSolveSingleElement(t *testing.T) {
	input := []int{42}
	solver := NewSolver(input)
	
	ops := solver.Solve()
	
	if len(ops) != 0 {
		t.Errorf("Expected 0 operations for single element, got %d", len(ops))
	}
}

func TestSolveTwoElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // expected number of operations
	}{
		{
			name:     "Already sorted",
			input:    []int{1, 2},
			expected: 0,
		},
		{
			name:     "Needs swap",
			input:    []int{2, 1},
			expected: 1,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			if len(ops) != tt.expected {
				t.Errorf("Expected %d operations, got %d", tt.expected, len(ops))
			}
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Error("Solution should result in sorted stack")
			}
		})
	}
}

func TestSolveThreeElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Case 3,2,1", []int{3, 2, 1}},
		{"Case 3,1,2", []int{3, 1, 2}},
		{"Case 2,3,1", []int{2, 3, 1}},
		{"Case 2,1,3", []int{2, 1, 3}},
		{"Case 1,3,2", []int{1, 3, 2}},
		{"Case 1,2,3", []int{1, 2, 3}}, // already sorted
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("Solution for %v should result in sorted stack", tt.input)
			}
			
			// For 3 elements, should not need more than 2 operations (except for already sorted)
			if !isSorted(tt.input) && len(ops) > 2 {
				t.Errorf("Solution for %v uses %d operations, expected <= 2", tt.input, len(ops))
			}
		})
	}
}

func TestSolveFiveElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Reverse sorted", []int{5, 4, 3, 2, 1}},
		{"Random order", []int{3, 1, 4, 2, 5}},
		{"Mostly sorted", []int{1, 3, 2, 4, 5}},
		{"Four elements", []int{4, 3, 2, 1}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.input)
			ops := solver.Solve()
			
			// Verify the result is sorted
			result := validateSolution(tt.input, ops)
			if !result {
				t.Errorf("Solution for %v should result in sorted stack", tt.input)
			}
			
			// For 5 elements, should be reasonably efficient
			if len(ops) > 12 {
				t.Errorf("Solution for %v uses %d operations, expected <= 12", tt.input, len(ops))
			}
		})
	}
}

func TestSolveLargeStack(t *testing.T) {
	// Test with a larger stack
	input := []int{10, 3, 7, 1, 9, 2, 8, 4, 6, 5}
	solver := NewSolver(input)
	ops := solver.Solve()
	
	// Verify the result is sorted
	result := validateSolution(input, ops)
	if !result {
		t.Error("Solution should result in sorted stack")
	}
	
	// Should complete in reasonable number of operations
	if len(ops) > 100 {
		t.Errorf("Solution uses %d operations, might be inefficient", len(ops))
	}
}

func TestFindMinPosition(t *testing.T) {
	solver := NewSolver([]int{3, 1, 4, 2, 5})
	
	minPos := solver.findMinPosition(solver.stackA)
	
	if minPos != 1 {
		t.Errorf("Expected min position 1, got %d", minPos)
	}
	
	// Test with empty stack
	emptyStack := stack.NewEmptyStack()
	solver.stackA = emptyStack
	minPos = solver.findMinPosition(solver.stackA)
	
	if minPos != -1 {
		t.Errorf("Expected -1 for empty stack, got %d", minPos)
	}
}

func TestFindMaxPosition(t *testing.T) {
	solver := NewSolver([]int{3, 1, 4, 2, 5})
	
	maxPos := solver.findMaxPosition(solver.stackA)
	
	if maxPos != 4 {
		t.Errorf("Expected max position 4, got %d", maxPos)
	}
}

func TestMoveToTopOptimized(t *testing.T) {
	solver := NewSolver([]int{1, 2, 3, 4, 5})
	
	// Test moving element at position 2 to top
	solver.moveToTopOptimized(solver.stackA, 2, true)
	
	top, _ := solver.stackA.Top()
	if top != 3 {
		t.Errorf("Expected top element to be 3 after moveToTopOptimized, got %d", top)
	}
	
	// Test with position 0 (should do nothing)
	solver = NewSolver([]int{1, 2, 3, 4, 5})
	originalTop, _ := solver.stackA.Top()
	solver.moveToTopOptimized(solver.stackA, 0, true)
	newTop, _ := solver.stackA.Top()
	
	if originalTop != newTop {
		t.Error("moveToTopOptimized with position 0 should not change the stack")
	}
}

func TestMoveToTopOptimizedWithReverseRotate(t *testing.T) {
	solver := NewSolver([]int{1, 2, 3, 4, 5})
	
	// Test moving element at position 4 (last) to top - should use reverse rotate
	solver.moveToTopOptimized(solver.stackA, 4, true)
	
	top, _ := solver.stackA.Top()
	if top != 5 {
		t.Errorf("Expected top element to be 5 after moveToTopOptimized, got %d", top)
	}
}

func TestExecuteAndRecord(t *testing.T) {
	solver := NewSolver([]int{2, 1})
	
	initialOpsCount := len(solver.operations)
	solver.executeAndRecord(operations.SA)
	
	if len(solver.operations) != initialOpsCount+1 {
		t.Error("executeAndRecord should add operation to the list")
	}
	
	if solver.operations[0] != operations.SA {
		t.Errorf("Expected first operation to be SA, got %v", solver.operations[0])
	}
	
	// Check that the operation was actually executed
	top, _ := solver.stackA.Top()
	if top != 1 {
		t.Errorf("Expected top element to be 1 after SA, got %d", top)
	}
}

func TestFindTargetForChunk(t *testing.T) {
	solver := NewSolver([]int{5, 1, 4, 2, 3})
	
	target := solver.findTargetForChunk(2)
	
	// Should return median value (3 in this case)
	if target != 3 {
		t.Errorf("Expected target to be 3 (median), got %d", target)
	}
	
	// Test with empty stack
	solver.stackA = stack.NewEmptyStack()
	target = solver.findTargetForChunk(1)
	
	if target != 0 {
		t.Errorf("Expected target to be 0 for empty stack, got %d", target)
	}
}

func TestFindPositionOfTarget(t *testing.T) {
	solver := NewSolver([]int{5, 1, 4, 2, 3})
	
	pos := solver.findPositionOfTarget(solver.stackA, 4)
	
	if pos != 2 {
		t.Errorf("Expected position of target 4 to be 2, got %d", pos)
	}
	
	// Test with target not in stack - should find closest smaller
	pos = solver.findPositionOfTarget(solver.stackA, 6) // larger than all
	if pos < 0 || pos >= solver.stackA.Size() {
		t.Errorf("Expected valid position for non-existent target, got %d", pos)
	}
}

// Helper function to validate that a solution actually sorts the input
func validateSolution(input []int, ops []operations.Operation) bool {
	if len(input) == 0 {
		return true
	}
	
	stackA := stack.NewStack(input)
	stackB := stack.NewEmptyStack()
	
	// Execute all operations
	for _, op := range ops {
		err := operations.ExecuteOperation(stackA, stackB, op)
		if err != nil {
			return false
		}
	}
	
	// Check if stack A is sorted and stack B is empty
	return stackA.IsSorted() && stackB.IsEmpty()
}

// Helper function to check if a slice is sorted
func isSorted(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

// Benchmark tests for performance evaluation
func BenchmarkSolveSmall(b *testing.B) {
	input := []int{3, 1, 4, 2, 5}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}

func BenchmarkSolveMedium(b *testing.B) {
	input := []int{10, 3, 7, 1, 9, 2, 8, 4, 6, 5, 15, 12, 11, 13, 14}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}

func TestSpecificCase(t *testing.T) {
	// Test the specific case mentioned: "2 1 3 6 5 8"
	input := []int{2, 1, 3, 6, 5, 8}
	solver := NewSolver(input)
	ops := solver.Solve()
	
	// Should be less than 9 operations for this case
	if len(ops) >= 9 {
		t.Errorf("Expected less than 9 operations for input %v, got %d operations", input, len(ops))
		t.Logf("Operations: %v", ops)
	}
	
	// Verify the result is sorted
	result := validateSolution(input, ops)
	if !result {
		t.Errorf("Solution for %v should result in sorted stack", input)
	}
}

func BenchmarkSolveLarge(b *testing.B) {
	input := make([]int, 100)
	for i := 0; i < 100; i++ {
		input[i] = 100 - i // Reverse sorted
	}
	
	for i := 0; i < b.N; i++ {
		solver := NewSolver(input)
		solver.Solve()
	}
}