package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}

	fmt.Printf("Problem 1: %v\n", problem1(data))
	fmt.Printf("Problem 2: %v\n", problem2(data))
}
func convertToSet(data []int) map[int]bool {
	dataSet := make(map[int]bool)
	for _, i := range data {
		dataSet[i] = true
	}
	return dataSet
}

func problem1(data []int) int {
	dataSet := convertToSet(data)
	for _, i := range data {
		_, ok := dataSet[2020-i]
		if ok {
			return i * (2020 - i)
		}
	}
	return -1
}

func problem2(data []int) int {
	dataSet := convertToSet(data)
	for i := range data {
		for j := range data[i:] {
			_, ok := dataSet[2020-data[i]-data[j]]
			if ok {
				return data[i] * data[j] * (2020 - data[i] - data[j])
			}
		}
	}
	return -1
}
