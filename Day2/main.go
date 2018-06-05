package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	file, err := os.Open("seq.txt")
	check(err)
	defer file.Close()

	bufScanner := bufio.NewScanner(file)
	bufScanner.Split(bufio.ScanLines)

	var sum = 0
	var s scanner.Scanner

	for bufScanner.Scan() {
		row := bufScanner.Text()
		getRowDiff(&sum, s, row)
	}

	fmt.Printf("The checksum is  %d", sum)
}

func getRowDiff(sum *int, s scanner.Scanner, row string) {
	s.Init(strings.NewReader(row))
	var min = math.MaxUint32
	var max = 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		val, err := strconv.Atoi(s.TokenText())
		check(err)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	*sum += max - min
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
