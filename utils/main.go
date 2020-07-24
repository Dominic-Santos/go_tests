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

func RandomNumber(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

func RandomChoise(a, b int) int {
	if RandomNumber(0, 1) == 0 {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}