package util

import (
	"math/rand"
	"strings"
	"time"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(len(ALPHABET))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := [3]string{"USD", "EUR", "JPY"}
	return currencies[rand.Intn(3)]
}

