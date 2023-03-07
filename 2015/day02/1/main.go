package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func paper(row ...string) int {
	l, _ := strconv.Atoi(row[0])
	w, _ := strconv.Atoi(row[1])
	h, _ := strconv.Atoi(row[2])
	wrapping_paper := 2*l*w + 2*w*h + 2*l*h
	slack := min(l*w, min(w*h, h*l))
	return wrapping_paper + slack
}

func min(a int, b int) int {
	var ret int = 0
	if a < b {
		ret = a
	} else {
		ret = b
	}
	return ret
}

func main() {
	input := puzzleInput()
	total := 0
	for _, line := range input[:len(input)-1] {
		row := strings.Split(line, "x")
		present := paper(row...)
		total += present
	}
	fmt.Println(total)
}
