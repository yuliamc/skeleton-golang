package math

import (
	"math/rand"
	"strconv"
	"time"
)

func Random(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func ConvertStringToUint(s string) (uint, error) {
	temp, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(temp), nil
}
