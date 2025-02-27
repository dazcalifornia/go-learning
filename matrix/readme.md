# Matrix Creator

A simple Go program that creates and displays a numbered matrix based on user-specified dimensions.

## Description

This program prompts the user to input the horizontal and vertical dimensions of a matrix. It then creates a two-dimensional slice (matrix) of the specified size and fills it with sequential numbers from 1 to (horizontal × vertical). Finally, it displays the matrix in two formats: as raw slices and as a formatted grid.

## Features

- Dynamic matrix creation based on user input
- Sequential numbering of matrix elements
- Modern Go idioms using `range` for iteration
- Two display formats (raw and formatted grid)

## How It Works

1. The program asks for horizontal (columns) and vertical (rows) dimensions
2. It creates a 2D slice with the specified dimensions
3. It populates the matrix with sequential numbers (1, 2, 3, ...)
4. It displays the matrix in two different formats

## Code Overview

```go
package main

import (
	"fmt"
)

func main() {
	var horizontal, vertical int
	var matrix [][]int

	num := 0

	fmt.Print("Enter horizontal size of matrix: ")
	fmt.Scan(&horizontal)
	fmt.Print("Enter vertical size of matrix: ")
	fmt.Scan(&vertical)

	// Initialize the matrix
	matrix = make([][]int, vertical)

	// Modern approach using range for row initialization and filling
	for v := range matrix {
		matrix[v] = make([]int, horizontal)
		for h := range matrix[v] {
			num++
			matrix[v][h] = num
		}
	}

	// Print the matrix using range
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Alternatively, for more formatted output:
	fmt.Println("\nFormatted matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%3d ", val) // Pad with spaces for alignment
		}
		fmt.Println()
	}
}
```

## Usage

1. Compile the program with `go build` or run it directly with `go run main.go`
2. Enter the horizontal size (number of columns) when prompted
3. Enter the vertical size (number of rows) when prompted
4. View the generated matrix

### Example

```
Enter horizontal size of matrix: 4
Enter vertical size of matrix: 3
[1 2 3 4]
[5 6 7 8]
[9 10 11 12]

Formatted matrix:
  1   2   3   4
  5   6   7   8
  9  10  11  12
```

## Key Go Concepts Demonstrated

- Two-dimensional slices
- Dynamic memory allocation with `make()`
- Modern iteration with `range`
- User input with `fmt.Scan()`
- Formatted output with `fmt.Printf()`

## Notes for Developers

- The matrix is filled row by row (row-major order)
- The first element starts at 1 and increments sequentially
- To modify the starting number or increment pattern, adjust the `num` variable and its modification

## Extensions

Possible enhancements to this program:

- Add options for different filling patterns (spirals, diagonals, etc.)
- Support for non-integer data types
- Matrix operations (addition, multiplication, transposition)
- File I/O for saving and loading matrices
