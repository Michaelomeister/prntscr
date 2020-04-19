package randomizer

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "0123456789"

//pointer seed based on time
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//random string of desired length
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//I don't even know why I separated this func
func String(length int) string {
	return StringWithCharset(length, charset)
}
