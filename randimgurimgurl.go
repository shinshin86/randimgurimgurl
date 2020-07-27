package randimgurimgurl

import (
	"math/rand"
	"time"
)

const rsLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Getter() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = rsLetters[rand.Intn(len(rsLetters))]
	}
	extArr := []string{".jpg", ".png"}
	ext := extArr[rand.Intn(len(extArr))]

	return "https://i.imgur.com/" + string(b) + ext
}
