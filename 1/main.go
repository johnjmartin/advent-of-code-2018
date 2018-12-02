package main

import (
	"fmt"
	"strconv"

	"github.com/johnjmartin/advent-of-code-2018/shared"
)

/**
 */

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	x := 0

	var l []int
	set := make(map[int]bool)

	set[x] = true
	for {

		fileLines := shared.FileLineReader("input.txt")

		for _, s := range fileLines {
			i, err := strconv.Atoi(s)
			check(err)
			x += i

			l = append(l, x)
			if _, ok := set[x]; ok {
				fmt.Println(x)
				return
			} else {
				set[x] = true
			}
		}

	}
}
