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

// Convert string to uint.
// @TODO Bener gak nih disini?
func ConvertStringToUint(s string) (uint, error) {
	temp, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(temp), nil
}
