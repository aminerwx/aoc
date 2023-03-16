package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func puzzleInput() []byte {
	_, file, _, _ := runtime.Caller(0)
	folder := filepath.Dir(filepath.Dir(file))
	dat, err := os.ReadFile(filepath.Join(folder, "input"))
	if err != nil {
		panic(err)
	}
	return dat
}

func main() {
	input := puzzleInput()
	lines := bytes.Split(input, []byte("\n"))
	strCode := 0
	strMemory := 0
	for _, s := range lines[:len(lines)-1] {
		strCode += len(s)
		v := strconv.Quote(string(s))
		strMemory += len(v)
	}
	fmt.Println(strMemory, "-", strCode, "=", strMemory-strCode)
}
