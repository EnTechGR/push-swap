package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := NewStack(data)
	
	if s.Size() != 5 {
		t.Errorf("Expected size 5, got %d", s.Size())
	}
	
	// Test that stack data is independent of original slice
	data[0] = 999
	top, _ := s.Top()
	if top == 999 {
		t.Error("Stack should be independent of original slice")
	}
}

func TestNewEmptyStack(t *testing.T) {
	s := NewEmptyStack()
	
	if !s.IsEmpty() {
		t.Error("New empty stack should be empty")
	}
	
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestPushPop(t *testing.T) {
	s := NewEmptyStack()
	
	// Test push
	s.Push(42)
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after push, got %d", s.Size())
	}
	
	top, err := s.Top()
	if err != nil {
		t.Errorf("Unexpected error getting top: %v", err)
	}
	if top != 42 {
		t.Errorf("Expected top 42, got %d", top)
	}
	
	// Test pop
	value, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error popping: %v", err)
	}
	if value != 42 {
		t.Errorf("Expected popped value 42, got %d", value)
	}
	
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping last element")
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewEmptyStack()
	
	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
}

func TestTopEmptyStack(t *testing.T) {
	s := NewEmptyStack()
	
	_, err := s.Top()
	if err == nil {
		t.Error("Expected error when getting top of empty stack")
	}
}

func TestAt(t *testing.T) {
	s := NewStack([]int{10, 20, 30})
	
	// Test valid indices
	val, err := s.At(0)
	if err != nil || val != 10 {
		t.Errorf("Expected At(0) to return 10, got %d (error: %v)", val, err)
	}
	
	val, err = s.At(2)
	if err != nil || val != 30 {
		t.Errorf("Expected At(2) to return 30, got %d (error: %v)", val, err)
	}
	
	// Test invalid indices
	_, err = s.At(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	_, err = s.At(3)
	if err == nil {
		t.Error("Expected error for out of bounds index")
	}
}

func TestIsSorted(t *testing.T) {
	// Test sorted stack
	sorted := NewStack([]int{1, 2, 3, 4, 5})
	if !sorted.IsSorted() {
		t.Error("Expected sorted stack to return true for IsSorted()")
	}
	
	// Test unsorted stack
	unsorted := NewStack([]int{3, 1, 4, 2, 5})
	if unsorted.IsSorted() {
		t.Error("Expected unsorted stack to return false for IsSorted()")
	}
	
	// Test single element
	single := NewStack([]int{42})
	if !single.IsSorted() {
		t.Error("Expected single element stack to be sorted")
	}
	
	// Test empty stack
	empty := NewEmptyStack()
	if !empty.IsSorted() {
		t.Error("Expected empty stack to be sorted")
	}
}

func TestToSlice(t *testing.T) {
	data := []int{1, 2, 3}
	s := NewStack(data)
	
	slice := s.ToSlice()
	
	// Test that slice has correct values
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
	
	for i, val := range slice {
		if val != data[i] {
			t.Errorf("Expected slice[%d] = %d, got %d", i, data[i], val)
		}
	}
	
	// Test that modifying returned slice doesn't affect stack
	slice[0] = 999
	top, _ := s.Top()
	if top == 999 {
		t.Error("Modifying returned slice should not affect stack")
	}
}

func TestClone(t *testing.T) {
	original := NewStack([]int{1, 2, 3})
	clone := original.Clone()
	
	// Test that clone has same values
	if clone.Size() != original.Size() {
		t.Error("Clone should have same size as original")
	}
	
	for i := 0; i < original.Size(); i++ {
		origVal, _ := original.At(i)
		cloneVal, _ := clone.At(i)
		if origVal != cloneVal {
			t.Errorf("Clone values should match original at index %d", i)
		}
	}
	
	// Test that clone is independent
	clone.Push(999)
	if original.Size() == clone.Size() {
		t.Error("Clone should be independent of original")
	}
}