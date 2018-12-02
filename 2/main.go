package main

import (
	"fmt"

	"github.com/johnjmartin/advent-of-code-2018/shared"
)

func main() {
	l := shared.FileLineReader("input.txt")
	partOne(l)
}

func partOne(l []string) {
	occurences := make(map[string]int)
	occurences["triples"] = 0
	occurences["doubles"] = 0

	for _, line := range l {
		countOccurences(occurences, line)
	}
	fmt.Println(occurences["triples"] * occurences["doubles"])
}

func countOccurences(occurences map[string]int, code string) {
	charMap := make(map[rune]int)
	for _, char := range code {
		charMap[char] += 1
	}
	for _, v := range charMap {
		if v == 3 {
			occurences["triples"] += 1
			break
		}
	}
	for _, v := range charMap {
		if v == 2 {
			occurences["doubles"] += 1
			break
		}
	}
}
