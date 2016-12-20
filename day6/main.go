package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cols []map[string]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cols == nil {
			cols = make([]map[string]int, len(line))
		}
		for i := range line {
			if col := cols[i]; col == nil {
				cols[i] = map[string]int{}
			}
			cols[i][string(line[i])]++
		}
	}

	var b bytes.Buffer
	for _, col := range cols {
		b.WriteString(minLetter(col))
	}
	fmt.Println(b.String())
}

func maxLetter(letters map[string]int) string {
	var maxLetter string
	var maxValue int

	for k, v := range letters {
		if v > maxValue {
			maxValue = v
			maxLetter = k
		}
	}

	return maxLetter
}

func minLetter(letters map[string]int) string {
	var minLetter string
	var minValue = int(^uint(0) >> 1)

	for k, v := range letters {
		if v < minValue {
			minValue = v
			minLetter = k
		}
	}

	return minLetter
}
