package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("^(([A-Za-z]+-)*)([0-9]+)\\[([A-Za-z]{5})\\]$")

func parseLine(line string) (words []string, sector int, checksum string) {
	parsed := re.FindStringSubmatch(line)
	if len(parsed) != 5 {
		panic(parsed)
	}
	words = strings.Split(parsed[1], "-")
	sector, err := strconv.Atoi(parsed[3])
	if err != nil {
		panic(err)
	}
	checksum = parsed[4]
	return
}

func makeChecksum(words []string) string {
	chars := make(map[string]int)

	for _, word := range words {
		for i := range word {
			chars[string(word[i])]++
		}
	}

	var letters []letter
	for value, frequency := range chars {
		l := letter{
			value:     value,
			frequency: frequency,
		}
		letters = append(letters, l)
	}

	sort.Sort(sort.Reverse(letterList(letters)))

	var vals []string
	for _, letter := range letters {
		vals = append(vals, letter.value)
	}

	return strings.Join(vals[:5], "")
}

func isValidLine(line string) (sector int, ok bool) {
	words, sector, checksum := parseLine(line)
	calculated := makeChecksum(words)
	if calculated != checksum {
		return 0, false
	}
	fmt.Println(shiftWords(words, sector%26), sector)
	return sector, true
}

const space = " "

func shiftWords(words []string, n int) string {
	var b bytes.Buffer

	for _, word := range words {
		for _, char := range word {
			b.WriteRune(shiftLetter(char, n))
		}
		b.WriteString(space)
	}
	return b.String()
}

func shiftLetter(char rune, n int) rune {
	shift, offset := rune(n), rune(26)

	if char+shift < 'z' {
		return char + shift
	}
	return char + shift - offset
}

type letter struct {
	value     string
	frequency int
}

type letterList []letter

func (l letterList) Len() int { return len(l) }
func (l letterList) Less(i, j int) bool {
	if l[i].frequency == l[j].frequency {
		return l[i].value > l[j].value
	}
	return l[i].frequency < l[j].frequency
}
func (l letterList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var c int
	for scanner.Scan() {
		line := scanner.Text()
		if sector, ok := isValidLine(line); ok {
			c += sector
		}
	}
	fmt.Println(c)
}
