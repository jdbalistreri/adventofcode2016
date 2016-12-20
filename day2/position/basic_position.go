package position

import (
	"strconv"
)

type position int

// NewBasicPosition ...
func NewBasicPosition() Position {
	p := position(5)
	return &p
}

func (p *position) Up() {
	if p == nil {
		return
	}
	if *p <= 3 {
		return
	}
	*p -= 3
}

func (p *position) Down() {
	if p == nil {
		return
	}
	if *p >= 7 {
		return
	}
	*p += 3
}

var leftBoundary = map[position]bool{
	1: true,
	4: true,
	7: true,
}

func (p *position) Left() {
	if p == nil {
		return
	}
	if _, ok := leftBoundary[*p]; ok {
		return
	}
	*p--
}

var rightBoundary = map[position]bool{
	3: true,
	6: true,
	9: true,
}

func (p *position) Right() {
	if p == nil {
		return
	}
	if _, ok := rightBoundary[*p]; ok {
		return
	}
	*p++
}

func (p *position) String() string {
	return strconv.Itoa(int(*p))
}
