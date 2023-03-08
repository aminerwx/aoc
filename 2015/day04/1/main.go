package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func lowestNumber(key string) int {
	for i := 0; true; i++ {
		mixer := key + strconv.Itoa(i)
		hash := md5.Sum([]byte(mixer))
		check := hex.EncodeToString(hash[:])
		if strings.HasPrefix(check, "00000") {
			return i
		}
	}
	return 0
}

func main() {
	input := "yzbqklnj"
	output := lowestNumber(input)
	fmt.Println("lowest number = ", output)
}
