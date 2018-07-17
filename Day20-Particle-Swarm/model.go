package main

import (
	"strconv"
)

type Particle struct {
	Pos *Position
	Vel *Velocity
	Acc *Acceleration
}

func NewParticle(pos []string, vel []string, acc []string) *Particle {
	return &Particle{
		Pos: newPosition(pos),
		Vel: newVelocity(vel),
		Acc: newAcceleration(acc),
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

func (p *Particle) Move() {
	// calculate velocity
	p.Vel.x += p.Acc.x
	p.Vel.y += p.Acc.y
	p.Vel.z += p.Acc.z
	// change position
	p.Pos.x += p.Vel.x
	p.Pos.y += p.Vel.y
	p.Pos.z += p.Vel.z
}

func (Particle) init() {

}
