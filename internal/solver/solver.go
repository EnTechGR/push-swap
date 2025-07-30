package solver

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// Solver represents the sorting algorithm solver
type Solver struct {
	stackA     *stack.Stack
	stackB     *stack.Stack
	operations []operations.Operation
}

// NewSolver creates a new solver with the given input
func NewSolver(input []int) *Solver {
	return &Solver{
		stackA:     stack.NewStack(input),
		stackB:     stack.NewEmptyStack(),
		operations: make([]operations.Operation, 0),
	}
}

// Solve solves the sorting problem and returns the operations
func (s *Solver) Solve() []operations.Operation {
	// If already sorted, return empty operations
	if s.stackA.IsSorted() {
		return s.operations
	}
	
	size := s.stackA.Size()
	
	switch {
	case size <= 1:
		// Already sorted or single element
		return s.operations
	case size == 2:
		s.solveTwo()
	case size == 3:
		s.solveThree()
	case size <= 5:
		s.solveFive()
	case size == 6:
		s.solveSix() // Specialized for 6 elements
	default:
		s.solveLarge()
	}
	
	return s.operations
}

// executeAndRecord executes an operation and records it
func (s *Solver) executeAndRecord(op operations.Operation) {
	operations.ExecuteOperation(s.stackA, s.stackB, op)
	s.operations = append(s.operations, op)
}

// solveTwo solves for 2 elements
func (s *Solver) solveTwo() {
	first, _ := s.stackA.At(0)
	second, _ := s.stackA.At(1)
	
	if first > second {
		s.executeAndRecord(operations.SA)
	}
}

// solveThree solves for 3 elements
func (s *Solver) solveThree() {
	for !s.stackA.IsSorted() {
		first, _ := s.stackA.At(0)
		second, _ := s.stackA.At(1)
		third, _ := s.stackA.At(2)
		
		if first > second && second < third && first < third {
			// Case: 2 1 3
			s.executeAndRecord(operations.SA)
		} else if first > second && second > third && first > third {
			// Case: 3 2 1
			s.executeAndRecord(operations.SA)
			s.executeAndRecord(operations.RRA)
		} else if first > second && second < third && first > third {
			// Case: 3 1 2
			s.executeAndRecord(operations.RA)
		} else if first < second && second > third && first < third {
			// Case: 1 3 2
			s.executeAndRecord(operations.SA)
			s.executeAndRecord(operations.RA)
		} else if first < second && second > third && first > third {
			// Case: 2 3 1
			s.executeAndRecord(operations.RRA)
		}
	}
}

// solveFive solves for 4-5 elements using optimized strategy
func (s *Solver) solveFive() {
	size := s.stackA.Size()
	
	// For 4-5 elements, move the 1-2 smallest elements to stack B
	elementsToMove := size - 3
	
	for i := 0; i < elementsToMove; i++ {
		minPos := s.findMinPosition(s.stackA)
		s.moveToTopOptimized(s.stackA, minPos, true)
		s.executeAndRecord(operations.PB)
	}
	
	// Sort remaining 3 elements in stack A
	s.solveThree()
	
	// Push back elements from stack B in correct order
	for !s.stackB.IsEmpty() {
		s.executeAndRecord(operations.PA)
	}
}

// solveLarge solves for large stacks using an optimized radix-like approach
func (s *Solver) solveLarge() {
	size := s.stackA.Size()
	
	// For stacks of 6 or more, use a more sophisticated approach
	if size <= 10 {
		s.solveSmallLarge()
	} else {
		s.solveLargeRadix()
	}
}

// solveSmallLarge handles stacks of 6-10 elements
func (s *Solver) solveSmallLarge() {
	// Push about half the elements to stack B, prioritizing smaller values
	targetSize := s.stackA.Size() / 2
	pushed := 0
	
	for pushed < targetSize && s.stackA.Size() > 3 {
		minPos := s.findMinPosition(s.stackA)
		s.moveToTopOptimized(s.stackA, minPos, true)
		s.executeAndRecord(operations.PB)
		pushed++
	}
	
	// Sort remaining elements in stack A
	s.solveThree()
	
	// Push back elements from stack B
	for !s.stackB.IsEmpty() {
		s.executeAndRecord(operations.PA)
	}
}

// solveLargeRadix uses a radix-sort inspired approach for large stacks
func (s *Solver) solveLargeRadix() {
	// Convert to ranks (0, 1, 2, ... n-1) for easier bit manipulation
	ranks := s.createRanks()
	s.applyRanks(ranks)
	
	// Sort by bits, starting from least significant
	maxBits := s.getMaxBits(len(ranks))
	
	for bit := 0; bit < maxBits; bit++ {
		size := s.stackA.Size()
		pushed := 0
		
		for i := 0; i < size; i++ {
			top, _ := s.stackA.Top()
			
			// If bit is 0, push to B; if bit is 1, rotate in A
			if (top>>bit)&1 == 0 {
				s.executeAndRecord(operations.PB)
				pushed++
			} else {
				s.executeAndRecord(operations.RA)
			}
		}
		
		// Push everything back from B to A
		for j := 0; j < pushed; j++ {
			s.executeAndRecord(operations.PA)
		}
	}
}

// findMinPosition finds the position of the minimum element in the stack
func (s *Solver) findMinPosition(st *stack.Stack) int {
	if st.IsEmpty() {
		return -1
	}
	
	minVal, _ := st.At(0)
	minPos := 0
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val < minVal {
			minVal = val
			minPos = i
		}
	}
	
	return minPos
}

// findMaxPosition finds the position of the maximum element in the stack
func (s *Solver) findMaxPosition(st *stack.Stack) int {
	if st.IsEmpty() {
		return -1
	}
	
	maxVal, _ := st.At(0)
	maxPos := 0
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val > maxVal {
			maxVal = val
			maxPos = i
		}
	}
	
	return maxPos
}

// moveToTopOptimized moves an element to top using the most efficient rotation
func (s *Solver) moveToTopOptimized(st *stack.Stack, position int, isStackA bool) {
	size := st.Size()
	if position == 0 || size <= 1 {
		return
	}
	
	// Choose most efficient direction
	if position <= size/2 {
		// Rotate forward (ra/rb)
		for i := 0; i < position; i++ {
			if isStackA {
				s.executeAndRecord(operations.RA)
			} else {
				s.executeAndRecord(operations.RB)
			}
		}
	} else {
		// Reverse rotate (rra/rrb)
		rotations := size - position
		for i := 0; i < rotations; i++ {
			if isStackA {
				s.executeAndRecord(operations.RRA)
			} else {
				s.executeAndRecord(operations.RRB)
			}
		}
	}
}

// createRanks creates a rank mapping for the input values
func (s *Solver) createRanks() map[int]int {
	data := s.stackA.ToSlice()
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	// Sort to create ranks
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	ranks := make(map[int]int)
	for i, val := range sorted {
		ranks[val] = i
	}
	
	return ranks
}

// applyRanks replaces stack values with their ranks
func (s *Solver) applyRanks(ranks map[int]int) {
	data := s.stackA.ToSlice()
	for i, val := range data {
		data[i] = ranks[val]
	}
	*s.stackA = *stack.NewStack(data)
}

// solveSix solves for exactly 6 elements using an optimized strategy
func (s *Solver) solveSix() {
	// For 6 elements, push the 2 smallest to stack B
	for i := 0; i < 2; i++ {
		minPos := s.findMinPosition(s.stackA)
		s.moveToTopOptimized(s.stackA, minPos, true)
		s.executeAndRecord(operations.PB)
	}
	
	// Now we have 4 elements in A, solve them by moving 1 more to B
	minPos := s.findMinPosition(s.stackA)
	s.moveToTopOptimized(s.stackA, minPos, true)
	s.executeAndRecord(operations.PB)
	
	// Sort remaining 3 elements in stack A
	s.solveThree()
	
	// Push back the 3 elements from stack B in order (they're already in order)
	s.executeAndRecord(operations.PA)
	s.executeAndRecord(operations.PA)
	s.executeAndRecord(operations.PA)
}

// findTargetForChunk finds a suitable target value for the current chunk
func (s *Solver) findTargetForChunk(chunkSize int) int {
	// Simple approach: find median-like value
	data := s.stackA.ToSlice()
	if len(data) == 0 {
		return 0
	}
	
	// Sort data to find median
	sortedData := make([]int, len(data))
	copy(sortedData, data)
	
	// Simple bubble sort for finding median
	for i := 0; i < len(sortedData)-1; i++ {
		for j := 0; j < len(sortedData)-i-1; j++ {
			if sortedData[j] > sortedData[j+1] {
				sortedData[j], sortedData[j+1] = sortedData[j+1], sortedData[j]
			}
		}
	}
	
	// Return median value
	return sortedData[len(sortedData)/2]
}

// findPositionOfTarget finds the position of the target value in the stack
func (s *Solver) findPositionOfTarget(st *stack.Stack, target int) int {
	for i := 0; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val == target {
			return i
		}
	}
	
	// If exact target not found, find closest smaller value
	closestPos := 0
	closestVal, _ := st.At(0)
	
	for i := 1; i < st.Size(); i++ {
		val, _ := st.At(i)
		if val < target && (val > closestVal || closestVal >= target) {
			closestVal = val
			closestPos = i
		}
	}
	
	return closestPos
}

// getMaxBits calculates the number of bits needed to represent n-1
func (s *Solver) getMaxBits(n int) int {
	if n <= 1 {
		return 1
	}
	bits := 0
	n--
	for n > 0 {
		n >>= 1
		bits++
	}
	return bits
}