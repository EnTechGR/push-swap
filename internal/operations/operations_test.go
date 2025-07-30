package operations

import (
	"push-swap/internal/stack"
	"testing"
)

func TestSwapA(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2, 3})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, SA)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	top, _ := stackA.Top()
	if top != 2 {
		t.Errorf("Expected top of stack A to be 2 after SA, got %d", top)
	}
	
	second, _ := stackA.At(1)
	if second != 1 {
		t.Errorf("Expected second element of stack A to be 1 after SA, got %d", second)
	}
}

func TestSwapAWithOneElement(t *testing.T) {
	stackA := stack.NewStack([]int{1})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, SA)
	if err != nil {
		t.Errorf("SA should not error with one element: %v", err)
	}
	
	top, _ := stackA.Top()
	if top != 1 {
		t.Errorf("Single element should remain unchanged after SA, got %d", top)
	}
}

func TestSwapB(t *testing.T) {
	stackA := stack.NewEmptyStack()
	stackB := stack.NewStack([]int{1, 2, 3})
	
	err := ExecuteOperation(stackA, stackB, SB)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	top, _ := stackB.Top()
	if top != 2 {
		t.Errorf("Expected top of stack B to be 2 after SB, got %d", top)
	}
}

func TestPushA(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2})
	stackB := stack.NewStack([]int{3, 4})
	
	err := ExecuteOperation(stackA, stackB, PA)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// Check that top of B was moved to top of A
	topA, _ := stackA.Top()
	if topA != 3 {
		t.Errorf("Expected top of stack A to be 3 after PA, got %d", topA)
	}
	
	// Check that B size decreased
	if stackB.Size() != 1 {
		t.Errorf("Expected stack B size to be 1 after PA, got %d", stackB.Size())
	}
	
	// Check that A size increased
	if stackA.Size() != 3 {
		t.Errorf("Expected stack A size to be 3 after PA, got %d", stackA.Size())
	}
}

func TestPushAFromEmptyB(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, PA)
	if err == nil {
		t.Error("Expected error when pushing from empty stack B")
	}
}

func TestPushB(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2})
	stackB := stack.NewStack([]int{3, 4})
	
	err := ExecuteOperation(stackA, stackB, PB)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// Check that top of A was moved to top of B
	topB, _ := stackB.Top()
	if topB != 1 {
		t.Errorf("Expected top of stack B to be 1 after PB, got %d", topB)
	}
	
	// Check sizes
	if stackA.Size() != 1 {
		t.Errorf("Expected stack A size to be 1 after PB, got %d", stackA.Size())
	}
	if stackB.Size() != 3 {
		t.Errorf("Expected stack B size to be 3 after PB, got %d", stackB.Size())
	}
}

func TestRotateA(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2, 3, 4})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, RA)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// After rotation: [2, 3, 4, 1]
	expected := []int{2, 3, 4, 1}
	for i, exp := range expected {
		val, _ := stackA.At(i)
		if val != exp {
			t.Errorf("Expected stackA[%d] = %d after RA, got %d", i, exp, val)
		}
	}
}

func TestReverseRotateA(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2, 3, 4})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, RRA)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// After reverse rotation: [4, 1, 2, 3]
	expected := []int{4, 1, 2, 3}
	for i, exp := range expected {
		val, _ := stackA.At(i)
		if val != exp {
			t.Errorf("Expected stackA[%d] = %d after RRA, got %d", i, exp, val)
		}
	}
}

func TestCombinedOperations(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2})
	stackB := stack.NewStack([]int{3, 4})
	
	// Test SS (swap both)
	err := ExecuteOperation(stackA, stackB, SS)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	topA, _ := stackA.Top()
	topB, _ := stackB.Top()
	
	if topA != 2 {
		t.Errorf("Expected top of A to be 2 after SS, got %d", topA)
	}
	if topB != 4 {
		t.Errorf("Expected top of B to be 4 after SS, got %d", topB)
	}
}

func TestExecuteOperations(t *testing.T) {
	stackA := stack.NewStack([]int{3, 2, 1})
	stackB := stack.NewEmptyStack()
	
	ops := []Operation{SA, PB, PB, SA, PA, PA}
	
	err := ExecuteOperations(stackA, stackB, ops)
	if err != nil {
		t.Errorf("Unexpected error executing operations: %v", err)
	}
	
	// Verify final state
	if !stackA.IsSorted() {
		t.Error("Stack A should be sorted after operations")
	}
	
	if !stackB.IsEmpty() {
		t.Error("Stack B should be empty after operations")
	}
}

func TestInvalidOperation(t *testing.T) {
	stackA := stack.NewStack([]int{1, 2, 3})
	stackB := stack.NewEmptyStack()
	
	err := ExecuteOperation(stackA, stackB, Operation("invalid"))
	if err == nil {
		t.Error("Expected error for invalid operation")
	}
}