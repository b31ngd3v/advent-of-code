package input

import (
	"log"
	"os"
)

func GetInput() string {
	input := os.Getenv("INPUT")
	file := "input.txt"
	if len(input) > 0 {
		file = input
	}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}
