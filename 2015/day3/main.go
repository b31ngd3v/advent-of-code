package main

import (
	"fmt"
	"strings"

	"github.com/b31ngd3v/advent-of-code/lib/input"
)

func main() {
	data := input.GetInput()
	partOne(data)
}

type HouseAddress struct {
	x int
	y int
}

func partOne(data string) {
	directions := strings.Split(data, "")
	x, y := 0, 0
	visitedHouses := []HouseAddress{{0, 0}}
	for _, direction := range directions {
		x, y = move(x, y, direction)
		curHouse := HouseAddress{x, y}
		if !isVisited(visitedHouses, curHouse) {
			visitedHouses = append(visitedHouses, HouseAddress{x, y})
		}
	}
	fmt.Println(len(visitedHouses))
}

func partTwo(data string) {
	directions := strings.Split(data, "")
	xSanta, ySanta := 0, 0
	xRoboSanta, yRoboSanta := 0, 0
	visitedHouses := []HouseAddress{{0, 0}}
	for i, direction := range directions {
		var x, y int
		if i%2 == 0 {
			xSanta, ySanta = move(xSanta, ySanta, direction)
			x, y = xSanta, ySanta
		} else {
			xRoboSanta, yRoboSanta = move(xRoboSanta, yRoboSanta, direction)
			x, y = xRoboSanta, yRoboSanta
		}
		curHouse := HouseAddress{x, y}
		if !isVisited(visitedHouses, curHouse) {
			visitedHouses = append(visitedHouses, HouseAddress{x, y})
		}
	}
	fmt.Println(len(visitedHouses))
}

func move(x, y int, direction string) (int, int) {
	switch direction {
	case "^":
		y++
	case "v":
		y--
	case ">":
		x++
	case "<":
		x--
	}
	return x, y
}

func isVisited(visitedHouses []HouseAddress, house HouseAddress) bool {
	for _, vh := range visitedHouses {
		if vh == house {
			return true
		}
	}
	return false
}
