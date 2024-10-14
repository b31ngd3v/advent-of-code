package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/b31ngd3v/advent-of-code/lib/input"
)

func main() {
	data := input.GetInput()
	partTwo(data)
}

func partOne(data string) {
	boxes := strings.Split(data, "\n")
	total := 0
	for _, box := range boxes {
		sides := make([]int, 3)
		for i, edge := range strings.Split(box, "x") {
			side, err := strconv.Atoi(edge)
			if err != nil {
				log.Fatalln(err)
			}
			sides[i] = side
		}
		sort.Ints(sides)
		total += sides[0]*sides[1]*2 + sides[1]*sides[2]*2 + sides[2]*sides[0]*2 + sides[0]*sides[1]
	}
	fmt.Println(total)
}

func partTwo(data string) {
	boxes := strings.Split(data, "\n")
	total := 0
	for _, box := range boxes {
		sides := make([]int, 3)
		for i, edge := range strings.Split(box, "x") {
			side, err := strconv.Atoi(edge)
			if err != nil {
				log.Fatalln(err)
			}
			sides[i] = side
		}
		sort.Ints(sides)
		total += (sides[0]+sides[1])*2 + sides[0]*sides[1]*sides[2]
	}
	fmt.Println(total)
}
