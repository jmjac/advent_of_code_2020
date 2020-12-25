package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	doorKey, _ := strconv.Atoi(strings.Split(string(data), "\n")[0])
	cardKey, _ := strconv.Atoi(strings.Split(string(data), "\n")[1])

	fmt.Printf("Part 1: %v\n", part1(doorKey, cardKey))
}

func part1(doorKey, cardKey int) int {
	cardLoopSize := findLoopSize(cardKey)
	privateKey := 1
	for i := 0; i < cardLoopSize; i++ {
		privateKey *= doorKey
		privateKey %= 20201227
	}

	return privateKey
}

func findLoopSize(key int) int {
	SN := 7
	transformed := 1
	for i := 1; i < 100000000000; i++ {
		transformed *= SN
		transformed %= 20201227
		if transformed == key {
			return i
		}
	}
	log.Panic("Coudn't find the loop size, increase the loop range and verify the input")
	return 0
}
