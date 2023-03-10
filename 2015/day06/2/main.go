package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func puzzleInput() []string {
	_, file, _, _ := runtime.Caller(0)
	folder := filepath.Dir(filepath.Dir(file))
	dat, err := os.ReadFile(filepath.Join(folder, "input"))
	lines := strings.Split(string(dat), "\n")
	if err != nil {
		panic(err)
	}
	return lines
}

type Coordinate struct {
	x, y int
}

type Instruction struct {
	state string
	src   Coordinate
	dest  Coordinate
}

const (
	OFF    = "turn off"
	ON     = "turn on"
	TOGGLE = "toggle"
)

func parser(input []string) []Instruction {
	statePattern := regexp.MustCompile(`(turn off|turn on|toggle)`)
	srcPattern := regexp.MustCompile(`([0-9]{1,3}.){2}`)
	destPattern := regexp.MustCompile(`([0-9]{1,3}(,)?){2}$`)
	var instructions []Instruction
	for _, line := range input {
		src := strings.Split(srcPattern.FindString(line), ",")
		dest := strings.Split(destPattern.FindString(line), ",")
		srcX, _ := strconv.Atoi(src[0])
		srcY, _ := strconv.Atoi(strings.TrimSpace(src[1]))
		destX, _ := strconv.Atoi(dest[0])
		destY, _ := strconv.Atoi(strings.TrimSpace(dest[1]))
		instructions = append(instructions, Instruction{
			statePattern.FindString(line),
			Coordinate{
				srcX,
				srcY,
			},
			Coordinate{
				destX,
				destY,
			},
		})
	}
	return instructions
}

func instruction(state string, src Coordinate, dest Coordinate, grid [1000][1000]int) [1000][1000]int {
	value := 0
	if state == OFF { value -= 1 }
	if state == ON  { value += 1 }
	if state == TOGGLE  { value += 2 }
	for i := src.x; i < dest.x+1; i++ {
		for j := src.y; j < dest.y+1; j++ {
			grid[i][j] += value
			if grid[i][j] < 0 {
				grid[i][j] = 0
			}
		}
	}
	return grid
}

func main() {
	var grid [1000][1000]int
	input := puzzleInput()
	instructions := parser(input)

	for _, ins := range instructions {
		grid = instruction(ins.state, ins.src, ins.dest, grid)
	}

	brightness := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			brightness += grid[i][j]
		}
	}
	fmt.Println(brightness)
}
