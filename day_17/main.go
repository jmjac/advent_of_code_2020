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

func countNear(x, y, z, w int, ignoreW bool, grid map[Point]int) int {
	count := 0
	for pX := x - 1; pX <= x+1; pX++ {
		for pY := y - 1; pY <= y+1; pY++ {
			for pZ := z - 1; pZ <= z+1; pZ++ {
				if ignoreW {
					count += grid[Point{pX, pY, pZ, 0}]
				} else {
					for pW := w - 1; pW <= w+1; pW++ {
						count += grid[Point{pX, pY, pZ, pW}]
					}
				}
			}
		}
	}
	return count - grid[Point{x, y, z, w}]
}
func changePoint(pX, pY, pZ, pW int, ignoreW bool, grid map[Point]int, newGrid map[Point]int) {
	count := countNear(pX, pY, pZ, pW, ignoreW, grid)
	if _, ok := grid[Point{pX, pY, pZ, pW}]; ok {
		if count == 2 || count == 3 {
			newGrid[Point{pX, pY, pZ, pW}] = 1
		}
	} else {
		if count == 3 {
			newGrid[Point{pX, pY, pZ, pW}] = 1
		}
	}
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
		for p := range grid {
			x, y, z, w := p.unpack()

			for pX := x - 1; pX <= x+1; pX++ {
				for pY := y - 1; pY <= y+1; pY++ {
					for pZ := z - 1; pZ <= z+1; pZ++ {
						if ignoreW {
							changePoint(pX, pY, pZ, 0, ignoreW, grid, newGrid)
						} else {
							for pW := w - 1; pW <= w+1; pW++ {
								changePoint(pX, pY, pZ, pW, ignoreW, grid, newGrid)
							}
						}

					}
				}
			}
		}
		grid = newGrid
	}
	return len(grid)
}
