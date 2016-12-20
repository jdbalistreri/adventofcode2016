package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
)

var doorID = "uqwqemis"

func hasFiveZeros(i int) (string, bool) {
	data := []byte(fmt.Sprintf("%s%d", doorID, i))
	v := fmt.Sprintf("%x", md5.Sum(data))
	return v[5:7], v[:5] == "00000"
}

func main() {
	total := map[int]string{}
	var i int
	for len(total) < 8 {
		i++
		if s, ok := hasFiveZeros(i); ok {
			pos, err := strconv.Atoi(string(s[0]))
			if err != nil {
				fmt.Println(err, pos, i)
				continue
			}
			if pos < 0 || pos > 7 {
				fmt.Println("invalid pos", pos)
				continue
			}
			if _, ok := total[pos]; !ok {
				total[pos] = string(s[1])
				fmt.Println(total)
			}
		}
	}

	var b bytes.Buffer
	for i = 0; i < 8; i++ {
		b.WriteString(total[i])
	}
	fmt.Println(b.String())
}
