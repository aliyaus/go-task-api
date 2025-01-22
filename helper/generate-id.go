package helper

import (
	"math/rand"
	"strconv"
)

// Utility function to generate a random ID
func GenerateID() string {
	return strconv.Itoa(rand.Intn(10000))
}
