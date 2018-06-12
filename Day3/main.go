package main

import (
	"fmt"
	"time"
)

type point struct {
	x int
	y int
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	values := make(chan *point, 2)
	start := time.Now().UTC()
	find := 368078
	// I am too lazy right now :)
	maxMatrixSize := 200000
	go onSpiral(0, maxMatrixSize/2, find, values)
	go onSpiral(maxMatrixSize/2, maxMatrixSize, find, values)

	one, another := <-values, <-values
	distance := Abs(one.x-another.x) + Abs(one.y-another.y)
	end := time.Now().UTC()
	fmt.Printf("The Manhattan Distance is: %d \n", distance)
	fmt.Printf("Took me %+v to find the answer.", end.Sub(start))
}

func onSpiral(from, to int, n int, vals chan *point) {
	for i := from; i < to; i++ {
		for j := from; j < to; j++ {
			x := min(min(i, j), min(n-1-i, n-1-j))
			var val int
			if i <= j {
				val = (n-2*x)*(n-2*x) - (i - x) - (j - x)
			} else {
				val = (n-2*x-2)*(n-2*x-2) + (i - x) + (j - x)
			}
			if val == n {
				vals <- &point{i, j}

			}
			if val == 1 {
				vals <- &point{i, j}

			}
		}
	}

}
