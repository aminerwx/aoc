package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func puzzle_input() []string {
    dat, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")
    return lines
}


func sort(a, b, c int) (int, int) {
    x, y := 0, 0
    if a <= b && a <= c {
        x = a 
        if b < c {
            y = b
        } else {
            y = c
        }
    }
    if b <= a && b <= c {
        x = b
        if a < c {
            y = a
        } else {
            y = c
        }
    }
    if c <= a && c <= b {
        x = c
        if b < a {
            y = b
        } else {
            y = a
        }
    }
    return x, y
}

func ribbon(row ...string) int {
    l, _ := strconv.Atoi(row[0])
    w, _ := strconv.Atoi(row[1])
    h, _ := strconv.Atoi(row[2])
    bow := l * w * h
    a, b := sort(l, w, h)
    return bow + a + a + b + b 
}

func main() {
   input := puzzle_input()
   total := 0
   for _, line := range input[:len(input) - 1] {
      row := strings.Split(line, "x")
      present := ribbon(row...)
      total += present
   } 
   fmt.Println(total)  
}
