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

func main() {
	input := puzzleInput()
	var floor int
	for i, char := range input {
		if char == 40 {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
