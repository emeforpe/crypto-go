package main

import (
	"cryptogo"
	"fmt"
)

func main() {

	alphabet := "abcdefghijklmnopqrstuwxyz"
	sentence := "thisisatest"

	key := cryptogo.VigenereKeyGenerator(alphabet, 3)
	cipherText := cryptogo.VigenereCipher(sentence, alphabet, key, "encrypt")
	plainText := cryptogo.VigenereCipher(cipherText, alphabet, key, "decrypt")

	fmt.Println(sentence)
	fmt.Println(cipherText)
	fmt.Println(plainText)
}
