package stack

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	data []int
}

// NewStack creates a new stack with the given data
func NewStack(data []int) *Stack {
	// Copy the slice to avoid external modifications
	stackData := make([]int, len(data))
	copy(stackData, data)
	return &Stack{data: stackData}
}

// NewEmptyStack creates a new empty stack
func NewEmptyStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Top returns the top element without removing it
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[0], nil
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value int) {
	s.data = append([]int{value}, s.data...)
}

// Pop removes and returns the top element
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	value := s.data[0]
	s.data = s.data[1:]
	return value, nil
}

// At returns the element at the given index (0 is top)
func (s *Stack) At(index int) (int, error) {
	if index < 0 || index >= len(s.data) {
		return 0, fmt.Errorf("index out of bounds")
	}
	return s.data[index], nil
}

// ToSlice returns a copy of the stack data
func (s *Stack) ToSlice() []int {
	result := make([]int, len(s.data))
	copy(result, s.data)
	return result
}

// IsSorted returns true if the stack is sorted in ascending order
func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] > s.data[i+1] {
			return false
		}
	}
	return true
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.data)
}

// Clone creates a deep copy of the stack
func (s *Stack) Clone() *Stack {
	return NewStack(s.data)
}