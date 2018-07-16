package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//particle swarm declaration
	//swarm := make([]Particle)

	f, err := os.Open("swarmcfg.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := scanner.Text()
		processRow(row)
	}

}

func processRow(row string) {

	particleData := strings.Split(row, ", ")

	pos := getPosition(particleData[0])
	fmt.Print(pos)

}

func getPosition(positionData string) []string {
	s := strings.Index(positionData, "<")

	s += len("<")
	e := strings.Index(positionData, ">")
	return strings.Split(positionData[s:e], ",")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
