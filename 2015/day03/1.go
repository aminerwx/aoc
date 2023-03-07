package main

import (
	"fmt"
	"os"
)

func puzzle_input() string {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return string(dat)
}

type Location struct {
	x, y int
}

func main() {
	input := puzzle_input()
	position := Location{}
	visited_house := []Location{
		Location{},
	}
	for _, current := range input[:len(input)-1] {
		visited := false
		switch current {
		case 94:
			position.y += 1
		case 62:
			position.x += 1
		case 118:
			position.y -= 1
		case 60:
			position.x -= 1
		}
		for j := range visited_house {
			if position.x == visited_house[j].x && position.y == visited_house[j].y {
				visited = true
				break
			}
		}
		if !visited {
			visited_house = append(visited_house, position)
		}
	}
	fmt.Println(len(visited_house))
}
