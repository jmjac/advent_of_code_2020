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

	player1 := make([]int, 0)
	player2 := make([]int, 0)

	lines := strings.Split(string(data), "\n\n")

	for _, i := range strings.Split(lines[0], "\n")[1:] {
		num, _ := strconv.Atoi(i)
		player1 = append(player1, num)
	}

	for _, i := range strings.Split(lines[1], "\n")[1:] {
		num, _ := strconv.Atoi(i)
		player2 = append(player2, num)
	}

	fmt.Printf("Part 1: %v\n", game(player1, player2))
	p1, p2 := recursiveGame(player1, player2)
	var winner int
	
	if p1 > p2 {
		winner = p1
	} else {
		winner = p2
	}
	fmt.Printf("Part 2: %v\n", winner)
}

func game(p1, p2 []int) int {
	for len(p1) != 0 && len(p2) != 0 {
		v1 := p1[0]
		p1 = p1[1:]
		v2 := p2[0]
		p2 = p2[1:]
		if v1 > v2 {
			p1 = append(p1, v1)
			p1 = append(p1, v2)
		} else {
			p2 = append(p2, v2)
			p2 = append(p2, v1)
		}
	}
	var winner []int
	if len(p1) != 0 {
		winner = p1
	} else {
		winner = p2
	}
	total := 0
	for i, v := range winner {
		total += (len(winner) - i) * v
	}
	return total
}

func playerHash(p []int) int64{
	var hash int64
	for i, v := range p{
		hash += int64((v+i)*i)
		hash *= 4139
		hash += int64(v)+int64(i)
		hash *= 5233
	}
	return hash
}

func recursiveGame(p1, p2 []int) (int, int) {
	history := make(map[int64]bool)
	for len(p1) != 0 && len(p2) != 0 {
		game := playerHash(p1)*7481 +  playerHash(p2)*7919
		if _, exists := history[game]; exists {
			break
		}
		history[game] = true
		v1 := p1[0]
		p1 = p1[1:]
		v2 := p2[0]
		p2 = p2[1:]
		if v1 > len(p1) || v2 > len(p2) {
			if v1 > v2 {
				p1 = append(p1, v1)
				p1 = append(p1, v2)
			} else {
				p2 = append(p2, v2)
				p2 = append(p2, v1)
			}
		} else {
			p1Slice := make([]int, v1)
			copy(p1Slice, p1[:v1])
			p2Slice := make([]int, v2)
			copy(p2Slice, p2[:v2])
			result, _ := recursiveGame(p1Slice[:], p2Slice[:])
			if result == 0 {
				p2 = append(p2, v2)
				p2 = append(p2, v1)
			} else {
				p1 = append(p1, v1)
				p1 = append(p1, v2)
			}
		}
	}
	total := 0
	for i, v := range p1 {
		total += (len(p1) - i) * v
	}
	total2 := 0
	for i, v := range p2 {
		total2 += (len(p1) - i) * v
	}
	return total, total2
}
