package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	boardingPasses := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		boardingPasses = append(boardingPasses, line)
	}
	fmt.Printf("Problem 1: %v\n", decodePasses(boardingPasses)[len(boardingPasses)-1])
	fmt.Printf("Problem 2: %v\n", problem2(boardingPasses))
}

func decodePasses(boardingPasses []string) []int {
	passes := make([]int, len(boardingPasses))
	for i, pass := range boardingPasses {
		sr, er := 0, 127
		sp, ep := 0, 7
		for _, k := range pass {
			switch k {
			case 'B':
				sr = (sr+er)/2 + 1
			case 'F':
				er = (sr + er) / 2
			case 'R':
				sp = (sp+ep)/2 + 1
			case 'L':
				ep = (sp + ep) / 2
			}
		}
		passes[i] = sr*8 + sp
	}
	sort.Ints(passes)
	return passes
}

func problem2(bordingPasses []string) int {
	intPasses := decodePasses(bordingPasses)
	last := intPasses[0]
	for _, i := range intPasses[1:] {
		if last+2 == i {
			return last + 1
		}
		last = i
	}
	return -1
}
