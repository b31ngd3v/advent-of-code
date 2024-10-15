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

func partTwo(data string) {
	strs := strings.Split(data, "\n")
	count := 0
	for _, str := range strs {
		if isSuperGoodString(str) {
			count++
		}
	}
	fmt.Println(count)
}

func isSuperGoodString(s string) bool {
	pairFound := false
	repeatingLtr := false
	for i := 0; i < len(s); i++ {
		if i+3 < len(s) && strings.Contains(s[i+2:], s[i:i+2]) {
			pairFound = true
		}
		if i+2 < len(s) && s[i] == s[i+2] {
			repeatingLtr = true
		}
	}
	return pairFound && repeatingLtr
}

func partOne(data string) {
	strs := strings.Split(data, "\n")
	count := 0
	for _, str := range strs {
		if isGoodString(str) {
			count++
		}
	}
	fmt.Println(count)
}

func isGoodString(s string) bool {
	vowelCount := 0
	repeatCh := false
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowelCount++
		}
		if len(s) != i+1 && s[i] == s[i+1] {
			repeatCh = true
		}
	}
	if repeatCh && vowelCount >= 3 && !isBadCombPresent(s) {
		return true
	}
	return false
}

func isBadCombPresent(s string) bool {
	badCombs := []string{"ab", "cd", "pq", "xy"}
	for _, badComb := range badCombs {
		if strings.Contains(s, badComb) {
			return true
		}
	}
	return false
}

func isVowel(ch byte) bool {
	vowels := "aeiou"
	return strings.ContainsRune(vowels, rune(ch))
}
