package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func hasZeroes(mixer string) bool {
	hash := md5.Sum([]byte(mixer))
	check := hex.EncodeToString(hash[:])
	return check[:5] == "00000"
}

func blender(key string) string {
	for i := 0; i < int(math.Pow10(len(key))); i++ {
		mixer := key + strconv.Itoa(i)
		found := hasZeroes(mixer)
		if found {
			return strconv.Itoa(i)
		}
	}
	return "not found"
}

func main() {
	input := "yzbqklnj"
	output := blender(input)
	fmt.Println("lowest number = ", output)
}
