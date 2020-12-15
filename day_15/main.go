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
	nums := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, i := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(i)
			nums = append(nums, num)
		}
	}

	fmt.Printf("Part 1: %v\n", game(nums, 2020))
	fmt.Printf("Part 2: %v\n", game(nums, 30000000))
}

func game(nums []int, endTurn int) int {
	mem := make(map[int]int)
	turn := 1
	num := 0
	for _, i := range nums {
		nextNum, ok := mem[i]
		if ok {
			num = nextNum - turn
		} else {
			num = 0
		}
		mem[i] = turn
		turn++
	}

	for turn != endTurn {
		nextNum, ok := mem[num]
		if ok {
			nextNum = turn - mem[num]
		}
		mem[num] = turn
		num = nextNum
		turn++
	}
	return num
}
