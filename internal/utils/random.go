package utils

import (
	"math/rand"
	"time"
)

// Seed initializes the random number generator
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// Intn returns a random integer in [0, n)
func Intn(n int) int {
	return rand.Intn(n)
}

// IntRange returns a random integer in [min, max]
func IntRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// String returns a random string from the provided slice
func String(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[rand.Intn(len(slice))]
}

// RandomString generates a random string of specified length from character set
func RandomString(chars string, length int) string {
	if len(chars) == 0 || length <= 0 {
		return ""
	}
	
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}