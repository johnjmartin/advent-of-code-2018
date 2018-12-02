package main

import (
	"fmt"

	"github.com/johnjmartin/advent-of-code-2018/shared"
)

func main() {
	lineList := shared.FileLineReader("input.txt")
	partOne(lineList)
	partTwo(lineList)
}

func partOne(lineList []string) {
	occurences := make(map[string]int)
	occurences["triplineListes"] = 0
	occurences["doublineListes"] = 0

	for _, line := range lineList {
		countOccurences(occurences, line)
	}
	fmt.Println(occurences["triplineListes"] * occurences["doublineListes"])
}

func countOccurences(occurences map[string]int, code string) {
	charMap := make(map[rune]int)
	for _, char := range code {
		charMap[char] += 1
	}
	for _, v := range charMap {
		if v == 3 {
			occurences["triplineListes"] += 1
			break
		}
	}
	for _, v := range charMap {
		if v == 2 {
			occurences["doublineListes"] += 1
			break
		}
	}
}

func partTwo(lineList []string) {
	for i, line := range lineList {
		checkForOneoff(line, lineList[i:])
	}
}

func checkForOneoff(line string, lineList []string) {
	for _, l := range lineList {
		differences := 0
		var index int
		for j, char := range line {
			if char != rune(l[j]) {
				differences += 1
				index = j
			}
		}
		if differences == 1 {
			fmt.Printf("Box id (%s) and (%s) have only one differing character index\n", line, l)
			fmt.Printf("Common Letters: %s \n", line[:index]+line[index+1:])
		}
	}
}
