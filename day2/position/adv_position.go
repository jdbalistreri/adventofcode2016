package position

import ()

type keypad struct {
	pad  [][]string
	x    int
	y    int
	maxX int
	maxY int
}

var SimplePad = [][]string{
	[]string{"0", "0", "0", "0", "0"},
	[]string{"0", "1", "2", "3", "0"},
	[]string{"0", "4", "5", "6", "0"},
	[]string{"0", "7", "8", "9", "0"},
	[]string{"0", "0", "0", "0", "0"},
}

var AdvPad = [][]string{
	[]string{"0", "0", "0", "0", "0", "0", "0"},
	[]string{"0", "0", "0", "1", "0", "0", "0"},
	[]string{"0", "0", "2", "3", "4", "0", "0"},
	[]string{"0", "5", "6", "7", "8", "9", "0"},
	[]string{"0", "0", "A", "B", "C", "0", "0"},
	[]string{"0", "0", "0", "D", "0", "0", "0"},
	[]string{"0", "0", "0", "0", "0", "0", "0"},
}

// NewAdvancedPosition ...
func NewAdvancedPosition(pad [][]string) Position {
	var k = keypad{
		pad:  pad,
		x:    2,
		y:    2,
		maxX: len(pad[0]),
		maxY: len(pad),
	}

	return &k
}

func (p *keypad) Up() {
	if p.y <= 0 {
		return
	}
	p.y--
	if p.String() == "0" {
		p.y++
	}
	return
}

func (p *keypad) Down() {
	if p.y >= p.maxY {
		return
	}
	p.y++
	if p.String() == "0" {
		p.y--
	}
	return
}

func (p *keypad) Left() {
	if p.x <= 0 {
		return
	}
	p.x--
	if p.String() == "0" {
		p.x++
	}
	return
}

func (p *keypad) Right() {
	if p.x >= p.maxX {
		return
	}
	p.x++
	if p.String() == "0" {
		p.x--
	}
	return
}

func (p *keypad) String() string {
	return p.pad[p.y][p.x]
}
