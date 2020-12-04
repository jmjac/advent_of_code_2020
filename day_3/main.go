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
	grid := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	fmt.Printf("Problem 1: %v\n", countTrees(grid, 3, 1))
	total := 1
	for i := 1; i <= 7; i += 2 {
		total *= countTrees(grid, i, 1)
	}
	total *= countTrees(grid, 1, 2)
	fmt.Printf("Problem 2: %v\n", total)
}

func countTrees(grid []string, stepRight int, stepDown int) int {
	count := 0
	length := len(grid[0])
	check := stepRight
	for i := stepDown; i < len(grid); i += stepDown {
		if grid[i][check%length] == '#' {
			count++
		}
		check += stepRight
	}
	return count
}
