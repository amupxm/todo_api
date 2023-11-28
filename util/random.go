package util

import (
	"fmt"
	"math/rand"
)

var alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(l int) string {
	letter := []rune(alphaNum)
	b := make([]rune, l)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomName() string {
	return RandomString(int(RandomInt(5, 10)))
}

func RandomOwner() string {
	return RandomString(int(RandomInt(5, 10)))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomName(), RandomName())
}
