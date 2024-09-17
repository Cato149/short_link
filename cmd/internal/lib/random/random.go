package random

import (
	"math/rand"
	"strings"
	"time"
)

func NewRandomName(lenght int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var sb strings.Builder

	for i := 0; i < lenght; i++ {
		sb.WriteRune(chars[rand.Intn(len(chars))])
	}

	return sb.String()
}
