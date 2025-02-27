package main

import (
	"fmt"
)

func main() {
	var horizontal, vertical int
	var matrix [][]int

	num := 0

	fmt.Print("Enter horizontal size of matrix:")
	fmt.Scan(&horizontal)
	fmt.Print("Enter vertical size of matrix:")
	fmt.Scan(&vertical)

	matrix = make([][]int, vertical)
	for v := range matrix {
		matrix[v] = make([]int, horizontal)
		for h := range matrix[v] {
			num++
			matrix[v][h] = num
		}
	}

	// Print the matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("\nFormatted matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%3d ", val) // Pad with spaces for alignment
		}
		fmt.Println()
	}
}
