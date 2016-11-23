package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := make([]byte, 32)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, plaintext)

	fmt.Println(out)
	fmt.Println(plaintext)
}
