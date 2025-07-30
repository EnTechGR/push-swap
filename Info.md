# Push-Swap Project Structure

```
push-swap/
├── cmd/
│   ├── push-swap/
│   │   └── main.go          # Main entry point for push-swap program
│   └── checker/
│       └── main.go          # Main entry point for checker program
├── internal/
│   ├── stack/
│   │   ├── stack.go         # Stack data structure and operations
│   │   └── stack_test.go    # Unit tests for stack operations
│   ├── parser/
│   │   ├── parser.go        # Input parsing and validation
│   │   └── parser_test.go   # Unit tests for parser
│   ├── solver/
│   │   ├── solver.go        # Sorting algorithm implementation
│   │   ├── small_sort.go    # Optimized sorting for small stacks
│   │   └── solver_test.go   # Unit tests for solver
│   └── operations/
│       ├── operations.go    # Operation definitions and execution
│       └── operations_test.go # Unit tests for operations
├── go.mod                   # Go module file
├── Makefile                 # Build automation
└── README.md               # Project documentation
```

## Implementation Plan

### 1. Stack Implementation
- Create a stack structure using Go slices
- Implement all 11 operations (sa, sb, ss, pa, pb, ra, rb, rr, rra, rrb, rrr)
- Add utility methods for stack manipulation

### 2. Input Parser
- Parse command-line arguments
- Validate integers and check for duplicates
- Handle edge cases (empty input, invalid formats)

### 3. Sorting Algorithm
- Implement different strategies based on stack size:
  - 2-3 elements: Direct optimal solutions
  - 4-5 elements: Small optimization algorithms
  - Larger stacks: More complex algorithms (possibly radix sort or chunk-based approach)

### 4. Operation Optimization
- Track and minimize operation count
- Implement operation combination logic (e.g., using rr instead of ra + rb)

### 5. Checker Implementation
- Read operations from stdin
- Execute operations on the input stack
- Verify final state (stack A sorted, stack B empty)