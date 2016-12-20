package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func legit(sides []int) bool {
	if len(sides) != 3 {
		fmt.Printf("illegal triangle w %d sides", len(sides))
		return false
	}
	sort.Ints(sides)
	return sides[0]+sides[1] > sides[2]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var c int
	scanner := bufio.NewScanner(file)
	buf := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, " ")
		intSides := []int{}
		for _, side := range sides {
			intSide, err := strconv.Atoi(side)
			if err != nil {
				continue
			}
			intSides = append(intSides, intSide)
		}
		buf = append(buf, intSides)

		if len(buf) == 3 {
			for i := 0; i < 3; i++ {
				var transposed []int
				for _, row := range buf {
					transposed = append(transposed, row[i])
				}
				fmt.Println(transposed, legit(transposed))
				if legit(transposed) {
					c++
				}
			}
			buf = [][]int{}
		}
	}
	fmt.Println(c)
}
