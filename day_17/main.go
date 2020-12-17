package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y, z, w int
}

func (p *Point) unpack() (int, int, int, int) {
	return p.x, p.y, p.z, p.w
}

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
	fmt.Printf("Part 1: %v\n", part1(data, 6, true))
	fmt.Printf("Part 2: %v\n", part1(data, 6, false))
}

func part1(data []string, cycles int, ignoreW bool) int {
	grid := make(map[Point]int)
	for x, line := range data {
		for y, r := range line {
			if r == '#' {
				grid[Point{x, y, 0, 0}] = 1
			}
		}
	}
	for i := 0; i < cycles; i++ {
		newGrid := make(map[Point]int)
		active := make(map[Point]int)
		for p := range grid {
			x, y, z, w := p.unpack()

			for pX := x - 1; pX <= x+1; pX++ {
				for pY := y - 1; pY <= y+1; pY++ {
					for pZ := z - 1; pZ <= z+1; pZ++ {
						if ignoreW {
							active[Point{pX, pY, pZ, 0}]++
						} else {
							for pW := w - 1; pW <= w+1; pW++ {
								active[Point{pX, pY, pZ, pW}]++
							}
						}

					}
				}
			}
			active[Point{x, y, z, w}]--

		}
		for p, act := range active {
			_, a := grid[p]
			if a && (act == 2 || act == 3) {
				newGrid[p] = 1
			} else if act == 3 {
				newGrid[p] = 1
			}
		}
		grid = newGrid
	}
	return len(grid)
}
