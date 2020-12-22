package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	player1 := make([]string, 0)
	player2 := make([]string, 0)

	lines := strings.Split(string(data), "\n\n")

	for _, i := range strings.Split(lines[0], "\n")[1:] {
		//num, _ := strconv.Atoi(i)
		player1 = append(player1, i)
	}

	for _, i := range strings.Split(lines[1], "\n")[1:] {
		//num, _ := strconv.Atoi(i)
		player2 = append(player2, i)
	}

	gamesHistory := make(map[string]bool)
	s := time.Now()
	fmt.Printf("Part 1: %v\n", game(player1, player2))
	fmt.Printf("Part 2: %v\n", recursiveGame(player1, player2, gamesHistory))
	fmt.Println(time.Since(s))
}

func game(p1, p2 []string) int {
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
	var winner []string
	if len(p1) != 0 {
		winner = p1
	} else {
		winner = p2
	}
	total := 0
	for i, v := range winner {
		num,_ := strconv.Atoi(v)
		total += (len(winner) - i) * num 
	}
	return total
}

func recursiveGame(p1, p2 []string, histcwory map[string]bool) int {
	history := make(map[string]bool)
	for len(p1) != 0 && len(p2) != 0 {
		game := p1[0]+","+p1[len(p1)-1]+","+p2[0]+","+p2[len(p2)-1] 
		if _, exists := history[game]; exists {
			break
		}
		history[game] = true
		n1 := p1[0]
		v1,_ := strconv.Atoi(n1)
		p1 = p1[1:]
		n2 := p2[0]
		v2,_ := strconv.Atoi(n2)
		p2 = p2[1:]
		if v1 > len(p1) || v2 > len(p2) {
			if v1 > v2 {
				p1 = append(p1, n1)
				p1 = append(p1, n2)
			} else {
				p2 = append(p2, n2)
				p2 = append(p2, n1)
			}
		} else {
			p1Slice := make([]string, v1)
			copy(p1Slice, p1[:v1])
			p2Slice := make([]string, v2)
			copy(p2Slice, p2[:v2])
			result := recursiveGame(p1Slice[:], p2Slice[:], history)
			if result == 0 {
				p2 = append(p2, n2)
				p2 = append(p2, n1)
			} else {
				p1 = append(p1, n1)
				p1 = append(p1, n2)
			}
		}
	}
	total := 0
	for i, v := range p1 {
		v, _ := strconv.Atoi(v)
		total += (len(p1) - i) * v
	}
	return total
}
