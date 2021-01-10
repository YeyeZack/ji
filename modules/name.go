package modules

import(
	"math/rand"
	"time"

)


// CreateName generates a random file name
func CreateName(extension string) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)+"."+extension
}
