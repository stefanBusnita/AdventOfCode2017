package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"sync"
)

func main() {

	//particle swarm declaration
	swarm := NewSwarm()

	f, err := os.Open("swarmcfg.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var particleNo = -1
	for scanner.Scan() {
		row := scanner.Text()
		particleNo++
		particle := getParticleFromRow(particleNo, row)
		//swarm = append(swarm, particle)
		swarm.addParticle(particle)
	}

	for i := 0; i < 10000; i++ {
		var wg sync.WaitGroup
		wg.Add(len(swarm))

		//move all particles one step
		for _, particle := range swarm {
			go particle.Move(&wg)
		}
		wg.Wait()

	}

	fmt.Printf("Closest is %d", swarm.findClosestToOrigin().Id)
}

func getParticleFromRow(particleNo int, row string) *Particle {

	particleData := strings.Split(row, ", ")

	posSlice := toStringSlice(particleData[0])
	velSlice := toStringSlice(particleData[1])
	accSlice := toStringSlice(particleData[2])

	return NewParticle(particleNo, posSlice, velSlice, accSlice)
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

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
