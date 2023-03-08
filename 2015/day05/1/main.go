package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

func hasVowels(input string) bool {
	counter := 0
	pattern := regexp.MustCompile(`[aieou]`)
	for i := 0; i < len(input); i++ {
		if pattern.MatchString(string(input[i])) {
			counter++
			if counter == 3 {
				return true
			}
		}
	}
	return false

}

func hasDuplicate(input string) bool {
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
			if input[i] == input[i+1] {
				return true
			}
		}
	}
	return false
}

func hasSubstring(input string) bool {
	return regexp.MustCompile(`ab|cd|pq|xy`).MatchString(string(input))
}

func main() {
	input := puzzleInput()
	counter := 0
	for _, line := range input {
		if hasVowels(string(line)) && hasDuplicate(string(line)) && !hasSubstring(string(line)) {
			counter++
		}
	}
	fmt.Println(counter)
}
