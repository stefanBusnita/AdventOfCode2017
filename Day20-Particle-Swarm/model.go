package main

type Particle struct {
	pos Position
	vel Velocity
	acc Acceleration
}

type Position struct {
	x int
	y int
	z int
}
type Velocity struct {
	x int
	y int
	z int
}
type Acceleration struct {
	x int
	y int
	z int
}

func (p *Particle) Move() {
	// calculate velocity
	p.vel.x += p.acc.x
	p.vel.y += p.acc.y
	p.vel.z += p.acc.z
	// change position
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func (Particle) init() {

}
