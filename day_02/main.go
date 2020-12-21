package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	sol1, sol2 := problem(data)
	fmt.Printf("Problem 1 & 2: %v & %v\n", sol1, sol2)
}

func problem(data []string) (int, int) {
	var validRanges int
	var validPositions int

	for _, line := range data {
		line := strings.Split(line, " ")
		lower, _ := strconv.Atoi(strings.Split(line[0], "-")[0])
		higher, _ := strconv.Atoi(strings.Split(line[0], "-")[1])
		required, _ := utf8.DecodeRuneInString(line[1])
		count := 0
		onlyOne := false

		for i, letter := range line[2] {
			if letter == required {
				count++
				if i+1 == lower || i+1 == higher {
					onlyOne = !onlyOne
				}
			}

		}

		if lower <= count && count <= higher {
			validRanges++
		}

		if onlyOne {
			validPositions++
		}

	}
	return validRanges, validPositions
}
