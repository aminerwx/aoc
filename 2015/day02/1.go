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


func paper(row ...string) int {
    l, _ := strconv.Atoi(row[0])
    w, _ := strconv.Atoi(row[1])
    h, _ := strconv.Atoi(row[2])
    wrapping_paper := 2 * l * w + 2 * w * h + 2 * l * h
    slack := min(l * w, min(w * h, h * l))
    return wrapping_paper + slack
}

func min(a int, b int) int {
    var ret int = 0
    if  a < b {
        ret = a
    } else {
        ret = b 
    }
    return ret
}

func main() {
   input := puzzle_input()
   total := 0
   for _, line := range input[:len(input) - 1] {
      row := strings.Split(line, "x")
      present := paper(row...)
      total += present
   } 
   fmt.Println(total)  
}
