package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lower       = "abcdefghijklmnopqrstuvwxyz"
	upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
	special     = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars    = lower + upper + digits + special
	passwordLen = 12 // hoặc để 8 nếu muốn ngắn
)

func init() {
	// Seed the random number generator with the current time
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func RandomFullname() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomPassword() string {
	var password strings.Builder

	// Đảm bảo có ít nhất 1 uppercase và 1 special
	password.WriteByte(upper[rand.Intn(len(upper))])
	password.WriteByte(special[rand.Intn(len(special))])

	// Thêm ngẫu nhiên các ký tự còn lại
	for password.Len() < passwordLen {
		password.WriteByte(allChars[rand.Intn(len(allChars))])
	}

	// Xáo trộn password
	runes := []rune(password.String())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return string(runes)
}
