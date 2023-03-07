package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func puzzleInput() string {
	_, file, _, _ := runtime.Caller(0)
	folder := filepath.Dir(filepath.Dir(file))
	dat, err := os.ReadFile(filepath.Join(folder, "input"))
	if err != nil {
		panic(err)
	}
	return string(dat)
}

type Location struct {
	x, y int
}

func isVisited(currentLocation Location, visitedHouses []Location) bool {
	visited := false
	for j := range visitedHouses {
		if currentLocation.x == visitedHouses[j].x && currentLocation.y == visitedHouses[j].y {
			visited = true
			break
		}
	}
	return visited
}
func delivery(input string) int {
	currentLocation := Location{}
	visitedHouses := []Location{
		{},
	}
	for _, direction := range input[:len(input)-1] {
		switch direction {
		case 94:
			currentLocation.y += 1
		case 62:
			currentLocation.x += 1
		case 118:
			currentLocation.y -= 1
		case 60:
			currentLocation.x -= 1
		}
		visited := isVisited(currentLocation, visitedHouses)
		if !visited {
			visitedHouses = append(visitedHouses, currentLocation)
		}
	}
	return len(visitedHouses)
}

func main() {
	input := puzzleInput()
	totalVisitedHouses := delivery(input)
	fmt.Println(totalVisitedHouses)
}
