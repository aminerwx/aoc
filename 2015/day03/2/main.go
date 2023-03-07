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

type Santa struct {
	name    string
	current Location
}

func direction(santa Santa, next int) Location {
	switch next {
	case 94:
		santa.current.y += 1
	case 62:
		santa.current.x += 1
	case 118:
		santa.current.y -= 1
	case 60:
		santa.current.x -= 1
	}
	return santa.current
}

func isVisited(currentLocation Location, visitedHouses []Location) []Location {
	visited := false
	for j := range visitedHouses {
		if currentLocation.x == visitedHouses[j].x && currentLocation.y == visitedHouses[j].y {
			visited = true
			break
		}
	}
	if !visited {
		visitedHouses = append(visitedHouses, currentLocation)
	}
	return visitedHouses
}

func delivery(input string) int {
	santa := Santa{name: "santa"}
	robot := Santa{name: "robot"}

	robotIndex := 1
	visitedHouses := []Location{{}}
	for i := 0; i < len(input)-1; i += 2 {
		santa.current = direction(santa, int(input[i]))

		if robotIndex < len(input) {
			robot.current = direction(robot, int(input[robotIndex]))
		}

		visitedHouses = isVisited(santa.current, visitedHouses)
		visitedHouses = isVisited(robot.current, visitedHouses)

		robotIndex += 2
	}
	return len(visitedHouses)
}

func main() {
	input := puzzleInput()
	totalVisitedHouses := delivery(input)
	fmt.Println(totalVisitedHouses)
}
