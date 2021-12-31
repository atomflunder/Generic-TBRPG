package utils

import (
	"math/rand"
	"time"
)

func NewRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomNumber(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}

func GetRandomNumberInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
