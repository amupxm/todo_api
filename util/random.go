package util

import (
	"math/rand"
)

var alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString() string {

	letter := []rune(alphaNum)
	b := make([]rune, 7)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
