package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}
	return i
}
