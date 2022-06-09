package random

import (
	"math/rand"
	"time"
)

const alphabet = "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// We need to seed on each package initialization to avoid repeating random codes
	rand.Seed(time.Now().UnixNano())
}

func String(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
