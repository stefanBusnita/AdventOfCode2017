package main

import (
	"strconv"
	"sync"
)

type Swarm map[int]*Particle

func NewSwarm() Swarm {
	return make(Swarm, 0)
}

func (s Swarm) addParticle(p *Particle) {
	s[p.Id] = p
}

func (s Swarm) removeParticle(id int) {
	delete(s, id)
}

func (s Swarm) findClosestToOrigin() Particle {

	var selkey = 0
	for key, _ := range s {
		selkey = key
		break
	}

	cp := *s[selkey]
	for _, particle := range s {
		if particle.Md <= cp.Md {
			cp = *particle
		}
	}
	return cp
}

type Particle struct {
	Id  int
	Pos Position
	Vel Velocity
	Acc Acceleration
	Md  int
}

func (p *Particle) String() string {
	return strconv.Itoa(p.Pos.x) + strconv.Itoa(p.Pos.y) + strconv.Itoa(p.Pos.z)
}

func NewParticle(id int, pos []string, vel []string, acc []string) *Particle {
	return &Particle{
		Id:  id,
		Pos: *newPosition(pos),
		Vel: *newVelocity(vel),
		Acc: *newAcceleration(acc),
	}
}

type Position struct {
	x int
	y int
	z int
}

func newPosition(posData []string) *Position {
	x, y, z := getCoordData(posData)
	return &Position{
		x, y, z,
	}
}

type Velocity struct {
	x int
	y int
	z int
}

func newVelocity(velData []string) *Velocity {
	x, y, z := getCoordData(velData)
	return &Velocity{
		x, y, z,
	}
}

type Acceleration struct {
	x int
	y int
	z int
}

func newAcceleration(accData []string) *Acceleration {
	x, y, z := getCoordData(accData)
	return &Acceleration{
		x, y, z,
	}
}

func getCoordData(posData []string) (int, int, int) {
	x, err := strconv.Atoi(posData[0])
	check(err)
	y, err := strconv.Atoi(posData[1])
	check(err)
	z, err := strconv.Atoi(posData[2])
	check(err)
	return x, y, z
}

func (p *Particle) Move(wg *sync.WaitGroup) {
	// calculate velocity
	p.Vel.x += p.Acc.x
	p.Vel.y += p.Acc.y
	p.Vel.z += p.Acc.z
	// change position
	p.Pos.x += p.Vel.x
	p.Pos.y += p.Vel.y
	p.Pos.z += p.Vel.z

	p.Md = Abs(p.Pos.x) + Abs(p.Pos.y) + Abs(p.Pos.z)

	wg.Done()
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
