package position

import (
	"fmt"
)

// Position is a position on a keypad
type Position interface {
	Up()
	Left()
	Right()
	Down()
	fmt.Stringer
}

// NextPosition ...
func NextPosition(p Position, directions string) string {
	for _, char := range directions {
		switch char {
		case 'U':
			p.Up()
		case 'L':
			p.Left()
		case 'R':
			p.Right()
		case 'D':
			p.Down()
		}
	}
	return p.String()
}
