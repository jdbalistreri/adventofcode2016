package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var finalInput = "L1, L3, L5, L3, R1, L4, L5, R1, R3, L5, R1, L3, L2, L3, R2, R2, L3, L3, R1, L2, R1, L3, L2, R4, R2, L5, R4, L5, R4, L2, R3, L2, R4, R1, L5, L4, R1, L2, R3, R1, R2, L4, R1, L2, R3, L2, L3, R5, L192, R4, L5, R4, L1, R4, L4, R2, L5, R45, L2, L5, R4, R5, L3, R5, R77, R2, R5, L5, R1, R4, L4, L4, R2, L4, L1, R191, R1, L1, L2, L2, L4, L3, R1, L3, R1, R5, R3, L1, L4, L2, L3, L1, L1, R5, L4, R1, L3, R1, L2, R1, R4, R5, L4, L2, R4, R5, L1, L2, R3, L4, R2, R2, R3, L2, L3, L5, R3, R1, L4, L3, R4, R2, R2, R2, R1, L4, R4, R1, R2, R1, L2, L2, R4, L1, L2, R3, L3, L5, L4, R4, L3, L1, L5, L3, L5, R5, L5, L4, L2, R1, L2, L4, L2, L4, L1, R4, R4, R5, R1, L4, R2, L4, L2, L4, R2, L4, L1, L2, R1, R4, R3, R2, R2, R5, L1, L2"
var testInput = "R8, R4, R4, R8"

type Direction int

const (
	Left Direction = iota
	Right
)

type Person struct {
	dx    int
	dy    int
	x     int
	y     int
	dupe  string
	dupeX int
	dupeY int
	seen  map[string]bool
}

func (p *Person) String() string {
	return fmt.Sprintf("I'm standing at (%d, %d) and facing (%d, %d).\n I am %d blocks away from where i started.\n The first dupe location i hit was (%d, %d) and it was %d blocks away",
		p.x, p.y, p.dx, p.dy,
		away(p.x, p.y),
		p.dupeX, p.dupeY,
		away(p.dupeX, p.dupeY))
}

func away(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (p *Person) Move(dir Direction, count int) {
	p.turn(dir)
	p.walk(count)
}

func (p *Person) walk(count int) {
	for i := 0; i < count; i++ {
		p.x += p.dx
		p.y += p.dy
		if p.dupe == "" {
			strDir := dstr(p.x, p.y)
			_, ok := p.seen[strDir]
			if ok {
				p.dupe = strDir
				p.dupeX = p.x
				p.dupeY = p.y
			}
			p.seen[strDir] = true
		}
	}
}

func (p *Person) turn(dir Direction) {
	var newDX, newDY int

	switch p.dx {
	case -1:
		newDX, newDY = 0, -1
	case 1:
		newDX, newDY = 0, 1
	default:
		switch p.dy {
		case -1:
			newDX, newDY = 1, 0
		case 1:
			newDX, newDY = -1, 0
		default:
			panic("bad input")
		}
	}
	if dir == Left {
		p.dx, p.dy = newDX, newDY
		return
	}
	p.dx, p.dy = newDX*-1, newDY*-1
}

func New() *Person {
	return &Person{
		dy:   1,
		seen: map[string]bool{},
	}
}

func dstr(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (p *Person) UseDirections(input string) {
	dirs := strings.Split(input, ", ")

	for _, v := range dirs {
		count, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}
		switch string(v[0]) {
		case "L":
			p.Move(Left, count)
		case "R":
			p.Move(Right, count)
		default:
			panic(fmt.Sprintf("invalid case %s", string(v[0])))
		}
	}
}

func main() {
	joe := New()
	joe.UseDirections(finalInput)
	fmt.Println(joe)
}
