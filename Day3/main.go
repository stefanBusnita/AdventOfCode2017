package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	printSpiral(10)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printSpiral(n int) {
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			// x stores the layer in which (i, j)th
			// element lies
			var x int

			// Finds minimum of four inputs
			x = min(min(i, j), min(n-1-i, n-1-j))

			// For upper right half
			if i <= j {
				fmt.Printf("%d\t ", (n-2*x)*(n-2*x)-(i-x)-(j-x))
			} else {
				fmt.Printf("%d\t ", (n-2*x-2)*(n-2*x-2)+(i-x)+(j-x))
			}
		}
		fmt.Print("\n")
	}
}
