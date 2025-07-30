package operations

import (
	"fmt"
	"push-swap/internal/stack"
)

// Operation represents a stack operation
type Operation string

const (
	SA  Operation = "sa"  // swap first 2 elements of stack a
	SB  Operation = "sb"  // swap first 2 elements of stack b
	SS  Operation = "ss"  // execute sa and sb
	PA  Operation = "pa"  // push top element of stack b to stack a
	PB  Operation = "pb"  // push top element of stack a to stack b
	RA  Operation = "ra"  // rotate stack a (shift up all elements by 1)
	RB  Operation = "rb"  // rotate stack b
	RR  Operation = "rr"  // execute ra and rb
	RRA Operation = "rra" // reverse rotate a (shift down all elements by 1)
	RRB Operation = "rrb" // reverse rotate b
	RRR Operation = "rrr" // execute rra and rrb
)

// ValidOperations contains all valid operations
var ValidOperations = map[string]Operation{
	"sa":  SA,
	"sb":  SB,
	"ss":  SS,
	"pa":  PA,
	"pb":  PB,
	"ra":  RA,
	"rb":  RB,
	"rr":  RR,
	"rra": RRA,
	"rrb": RRB,
	"rrr": RRR,
}

// ExecuteOperation executes a single operation on the given stacks
func ExecuteOperation(stackA, stackB *stack.Stack, op Operation) error {
	switch op {
	case SA:
		return swapA(stackA)
	case SB:
		return swapB(stackB)
	case SS:
		swapA(stackA) // Ignore errors for ss operation
		return swapB(stackB)
	case PA:
		return pushA(stackA, stackB)
	case PB:
		return pushB(stackA, stackB)
	case RA:
		return rotateA(stackA)
	case RB:
		return rotateB(stackB)
	case RR:
		rotateA(stackA) // Ignore errors for rr operation
		return rotateB(stackB)
	case RRA:
		return reverseRotateA(stackA)
	case RRB:
		return reverseRotateB(stackB)
	case RRR:
		reverseRotateA(stackA) // Ignore errors for rrr operation
		return reverseRotateB(stackB)
	default:
		return fmt.Errorf("unknown operation: %s", op)
	}
}

// ExecuteOperations executes a sequence of operations
func ExecuteOperations(stackA, stackB *stack.Stack, operations []Operation) error {
	for _, op := range operations {
		if err := ExecuteOperation(stackA, stackB, op); err != nil {
			return err
		}
	}
	return nil
}

// swapA swaps the first 2 elements of stack a
func swapA(stackA *stack.Stack) error {
	if stackA.Size() < 2 {
		return nil // No error, just no operation needed
	}
	
	first, _ := stackA.Pop()
	second, _ := stackA.Pop()
	
	stackA.Push(first)
	stackA.Push(second)
	
	return nil
}

// swapB swaps the first 2 elements of stack b
func swapB(stackB *stack.Stack) error {
	if stackB.Size() < 2 {
		return nil // No error, just no operation needed
	}
	
	first, _ := stackB.Pop()
	second, _ := stackB.Pop()
	
	stackB.Push(first)
	stackB.Push(second)
	
	return nil
}

// pushA pushes the top element of stack b to stack a
func pushA(stackA, stackB *stack.Stack) error {
	if stackB.IsEmpty() {
		return fmt.Errorf("cannot push from empty stack b")
	}
	
	value, err := stackB.Pop()
	if err != nil {
		return err
	}
	
	stackA.Push(value)
	return nil
}

// pushB pushes the top element of stack a to stack b
func pushB(stackA, stackB *stack.Stack) error {
	if stackA.IsEmpty() {
		return fmt.Errorf("cannot push from empty stack a")
	}
	
	value, err := stackA.Pop()
	if err != nil {
		return err
	}
	
	stackB.Push(value)
	return nil
}

// rotateA rotates stack a (shift up all elements by 1)
func rotateA(stackA *stack.Stack) error {
	if stackA.Size() < 2 {
		return nil // No operation needed
	}
	
	data := stackA.ToSlice()
	first := data[0]
	
	// Shift all elements up
	for i := 0; i < len(data)-1; i++ {
		data[i] = data[i+1]
	}
	data[len(data)-1] = first
	
	*stackA = *stack.NewStack(data)
	return nil
}

// rotateB rotates stack b
func rotateB(stackB *stack.Stack) error {
	if stackB.Size() < 2 {
		return nil // No operation needed
	}
	
	data := stackB.ToSlice()
	first := data[0]
	
	// Shift all elements up
	for i := 0; i < len(data)-1; i++ {
		data[i] = data[i+1]
	}
	data[len(data)-1] = first
	
	*stackB = *stack.NewStack(data)
	return nil
}

// reverseRotateA reverse rotates stack a (shift down all elements by 1)
func reverseRotateA(stackA *stack.Stack) error {
	if stackA.Size() < 2 {
		return nil // No operation needed
	}
	
	data := stackA.ToSlice()
	last := data[len(data)-1]
	
	// Shift all elements down
	for i := len(data) - 1; i > 0; i-- {
		data[i] = data[i-1]
	}
	data[0] = last
	
	*stackA = *stack.NewStack(data)
	return nil
}

// reverseRotateB reverse rotates stack b
func reverseRotateB(stackB *stack.Stack) error {
	if stackB.Size() < 2 {
		return nil // No operation needed
	}
	
	data := stackB.ToSlice()
	last := data[len(data)-1]
	
	// Shift all elements down
	for i := len(data) - 1; i > 0; i-- {
		data[i] = data[i-1]
	}
	data[0] = last
	
	*stackB = *stack.NewStack(data)
	return nil
}
