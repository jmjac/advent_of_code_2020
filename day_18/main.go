package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	part1, part2 := sumCalculations(data)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func sumCalculations(data []string) (int, int) {
	totalPart1 := 0
	totalPart2 := 0
	for _, line := range data {
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		totalPart1 += calculate(strings.Split(line, " "))
		line = prepareOrder(line)
		totalPart2 += calculate(strings.Split(line, " "))
	}
	return totalPart1, totalPart2
}

func prepareOrder(line string) string {
	newLine := "("
	for _, i := range strings.Split(line, " ") {
		if i == "(" || i == ")" {
			newLine += " " + i
			newLine += " " + i
		} else if i == "*" {
			newLine += " )"
			newLine += " *"
			newLine += " ("
		} else if i != "+" {
			newLine += " ("
			newLine += " " + i
			newLine += " )"
		} else {
			newLine += " +"
		}
	}
	newLine += " )"
	return newLine
}

func evalute(data []string) int {
	s := make([]int, 0)
	for len(data) != 0 {
		t := data[0]
		data = data[1:]
		if t == "+" || t == "*" {
			x := s[len(s)-1]
			y := s[len(s)-2]
			s = s[:len(s)-2]
			if t == "+" {
				s = append(s, x+y)
			} else {
				s = append(s, x*y)
			}
		} else {
			n, err := strconv.Atoi(t)
			if t != "" && err != nil {
				log.Panic(err)
			}
			s = append(s, n)
		}
	}
	for _, i := range s {
		if i != 0 {
			return i
		}
	}
	return 0
}

func toPolishNotation(line []string) []string {
	operators := make([]string, len(line)/2)
	values := make([]string, len(line)/2)
	for i, _ := range line {
		i := line[len(line)-i-1]
		if i == "+" || i == "*" || i == ")" {
			operators = append(operators, i)
		} else if i == "(" {
			j := operators[len(operators)-1]
			operators = operators[:len(operators)-1]
			for j != ")" {
				values = append(values, j)
				j = operators[len(operators)-1]
				operators = operators[:len(operators)-1]
			}
		} else {
			values = append(values, i)
		}
	}

	for len(operators) != 0 {
		j := operators[len(operators)-1]
		values = append(values, j)
		operators = operators[:len(operators)-1]
	}

	return values
}

func calculate(line []string) int {
	values := toPolishNotation(line)
	return evalute(values)
}
