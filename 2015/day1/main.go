package main

import (
	"fmt"
	"strings"

	"github.com/b31ngd3v/advent-of-code/lib/input"
)

func main() {
	data := input.GetInput()
	partTwo(data)
}

func partOne(data string) {
	up := strings.Count(data, "(")
	down := strings.Count(data, ")")
	fmt.Println(up - down)
}

func partTwo(data string) {
	chars := strings.Split(data, "")
	level := 0
	for i, ch := range chars {
		switch ch {
		case "(":
			level++
		case ")":
			level--
		}
		if level == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
