package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//particle swarm declaration
	swarm := make([]*Particle, 0)

	f, err := os.Open("swarmcfg.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := scanner.Text()
		particle := processRow(row)
		swarm = append(swarm, particle)
	}

	for index, particle := range swarm {
		fmt.Printf("Currently at particle %d, pos: %+v, vel: %+v, acc:%+v \n", index, particle.Pos, particle.Vel, particle.Acc)
	}

}

func processRow(row string) *Particle {

	particleData := strings.Split(row, ", ")

	posSlice := toStringSlice(particleData[0])
	velSlice := toStringSlice(particleData[1])
	accSlice := toStringSlice(particleData[2])

	return NewParticle(posSlice, velSlice, accSlice)
}

func toStringSlice(positionData string) []string {
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
