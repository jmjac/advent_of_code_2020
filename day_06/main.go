package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	anserws := make([][]string, 0)
	group := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			anserws = append(anserws, group)
			group = make([]string, 0)
		} else {
			group = append(group, line)
		}
	}
	anserws = append(anserws, group)
	fmt.Printf("Problem 1: %v\n", count(anserws))
	fmt.Printf("Problem 2: %v\n", countAll(anserws))
}

func count(anserws [][]string) (total int) {
	for _, group := range anserws {
		count := 0
		setAns := make(map[rune]bool)
		for _, person := range group {
			for _, ans := range person {
				if _, ok := setAns[ans]; !ok {
					setAns[ans] = true
					count++
				}
			}
		}
		total += count
	}
	return
}

func countAll(anserws [][]string) (total int) {
	for _, group := range anserws {
		setAns := make(map[rune]int)
		for _, person := range group {
			for _, ans := range person {
				setAns[ans]++
			}
		}
		for _, freq := range setAns {
			if freq == len(group) {
				total++
			}
		}
	}
	return
}
