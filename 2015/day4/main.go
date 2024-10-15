package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/b31ngd3v/advent-of-code/lib/input"
)

func main() {
	data := input.GetInput()
	partTwo(data)
}

func partOne(data string) {
	a := findAns(data, 5)
	fmt.Println(a)
}

func partTwo(data string) {
	a := findAns(data, 6)
	fmt.Println(a)
}

func findAns(data string, n int) int {
	d := 0
	a := strings.Repeat("0", n)
	for {
		in := fmt.Sprintf("%s%d", data, d)
		hash := md5.Sum([]byte(in))
		res := hex.EncodeToString(hash[:])
		if strings.HasPrefix(res, a) {
			return d
		}
		d++
	}
}
