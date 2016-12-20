package main

import (
	"bufio"
	"fmt"
	"os"

	"code.uber.internal/rds/adventofcode/day2/position"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("basic_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var pos position.Position = position.NewBasicPosition()
	var pos position.Position = position.NewAdvancedPosition(position.AdvPad)
	results := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, position.NextPosition(pos, scanner.Text()))
	}
	fmt.Println(results)

	// file.Read
	// fmt.Println(p.nextPosition("ULL"))
	// fmt.Println(p.nextPosition("RRDDD"))
	// fmt.Println(p.nextPosition("LURDL"))
	// fmt.Println(p.nextPosition("UUUUD"))
}
