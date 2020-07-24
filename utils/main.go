package utils

import (
	"math/rand"
	"time"
)

func IntToChar(i int) string {
	return string('A' - 1 + i)
}

func CharToInt(s string) int {
	return int([]rune(s)[0])
}

func RandomNumber(min int, max int) int{
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}