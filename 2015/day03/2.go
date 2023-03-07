package main

import (
    "fmt"
    "os"
    "strings"
)

func puzzle_input() []string {
    dat, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")
    return lines
}

func main() {
   input := puzzle_input()

}
