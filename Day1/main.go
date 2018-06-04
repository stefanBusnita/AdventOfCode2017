package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	noCapcha    = "This is too complicated even for a machine... adkasjn unable to asndjkasnjdknjd.. help me... exit"
	capchaFound = "We found it sir ! Please try the following: %d"
)

func main() {

	input := retrieveInput("sequence.txt")

	if len(input) == 1 {
		log.Print("The input list should have at least 2 elements")
		log.Printf(noCapcha)
		os.Exit(1)
	}

	if len(input) == 2 {
		log.Print("Only 2 elements...")
		if input[0] == input[1] {
			log.Printf(capchaFound, input[0]+input[1])
		} else {
			log.Printf(noCapcha)
		}
		os.Exit(1)
	}

	var tmpSum = 0
	for i, v := range input {
		if i == len(input)-1 && v == input[0] {
			tmpSum += v
			continue
		}
		if v == input[i+1] {
			tmpSum += v
			continue
		}
	}

	log.Printf(capchaFound, tmpSum)
}

func retrieveInput(path string) []int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var input []int

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		input = append(input, val)
	}
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
