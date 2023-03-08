package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func puzzleInput() []string {
	_, file, _, _ := runtime.Caller(0)
	folder := filepath.Dir(filepath.Dir(file))
	dat, err := os.ReadFile(filepath.Join(folder, "input"))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	return lines
}

func hasDoublePair(input string) bool {
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
			a, b := input[i], input[i+1]
			for j := i + 2; j < len(input); j++ {
				if j+1 < len(input) {
					c, d := input[j], input[j+1]
					if a == c && b == d {
						return true
					}
				}
			}
		}
	}
	return false
}

func isRepeatable(input string) bool {
	for i := 0; i < len(input); i++ {
		if i+2 < len(input) {
			if input[i] == input[i+2] {
				return true
			}
		}
	}
	return false
}

func main() {
	input := puzzleInput()
	counter := 0
	for _, line := range input {
		if hasDoublePair(line) && isRepeatable(line) {
			counter++
		}
	}
	fmt.Println(counter)
}
