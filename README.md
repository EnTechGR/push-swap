# Push-Swap

A Go implementation of the classic push-swap sorting algorithm project. This project implements a non-comparative sorting algorithm using two stacks and a specific set of operations.

## Project Structure

```
push-swap/
├── cmd/
│   ├── push-swap/          # Main push-swap program
│   └── checker/            # Checker program for validation
├── internal/
│   ├── stack/              # Stack data structure implementation
│   ├── operations/         # Stack operations (sa, sb, pa, pb, etc.)
│   ├── parser/            # Input parsing and validation
│   └── solver/            # Sorting algorithm implementation
├── go.mod                 # Go module file
├── Makefile              # Build automation
└── README.md             # This file
```

## Programs

### push-swap
Calculates and displays the smallest program using push-swap instruction language that sorts the given integer arguments.

### checker
Takes integer arguments and reads instructions from standard input. Executes the instructions and displays `OK` if the integers are sorted correctly, otherwise displays `KO`.

## Available Operations

- `sa` - swap first 2 elements of stack a
- `sb` - swap first 2 elements of stack b  
- `ss` - execute sa and sb
- `pa` - push top element of stack b to stack a
- `pb` - push top element of stack a to stack b
- `ra` - rotate stack a (shift up all elements by 1)
- `rb` - rotate stack b
- `rr` - execute ra and rb
- `rra` - reverse rotate a (shift down all elements by 1)
- `rrb` - reverse rotate b
- `rrr` - execute rra and rrb

## Building

### Using Make
```bash
# Build both programs
make build

# Build individual programs
make push-swap
make checker

# Run tests
make test

# Format and vet code
make check

# Clean binaries
make clean
```

### Using Go directly
```bash
# Build push-swap
go build -o push-swap ./cmd/push-swap

# Build checker
go build -o checker ./cmd/checker
```

## Usage

### push-swap
```bash
# Basic usage
./push-swap "4 67 3 87 23"

# Multiple arguments
./push-swap 4 67 3 87 23

# Count operations
./push-swap "4 67 3 87 23" | wc -l
```

### checker
```bash
# Validate push-swap output
./push-swap "4 67 3 87 23" | ./checker "4 67 3 87 23"

# Manual input
echo -e "sa\npb\npa" | ./checker "3 2 1"

# Interactive mode
./checker "3 2 1"
sa
rra
pb
^D
```

## Examples

```bash
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa

$ ./push-swap "0 one 2 3"
Error

$ ./push-swap
(no output)

$ ./checker "3 2 1 0"
sa
rra
pb
KO

$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

## Algorithm Strategy

The solver uses different strategies based on stack size:

- **2-3 elements**: Direct optimal solutions
- **4-5 elements**: Optimized small sorting algorithms
- **Larger stacks**: Chunk-based approach with strategic element positioning

## Error Handling

The programs handle various error conditions:
- Invalid integers in input
- Duplicate numbers
- Invalid operations (checker only)
- Empty stacks during operations

Errors are displayed as "Error" followed by a newline on stderr.

## Testing

Run the test suite:
```bash
make test
```

Individual package tests:
```bash
go test ./internal/stack
go test ./internal/operations
go test ./internal/parser
```

## Development

### Code Quality
```bash
# Format code
make fmt

# Run static analysis
make vet

# Run all checks
make check
```

### Demo and Validation
```bash
# Run demo with example
make demo

# Validate push-swap output
make validate
```

## Requirements

- Go 1.21 or later
- Standard Go packages only (no external dependencies)

## License

This project is for educational purposes as part of the push-swap algorithm challenge.