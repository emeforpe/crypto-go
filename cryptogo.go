package cryptogo

import (
	"math/rand"
	"strings"
	"time"
)

func VigenereKeyGenerator(alphabet string, length int) string {

	var key []string

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	for len(key) < length {
		index := rnd.Intn(length)
		element := string(alphabet[index])
		key = append(key, element)
	}

	return strings.Join(key[:], "")
}

func VigenereCipher(text string, alphabet string, key string, mode string) string {

	keyLenght := len(key)
	alphabetChars := strings.Split(alphabet, "")
	keyChars := strings.Split(key, "")
	count := 0
	result := ""

	for _, letter := range strings.Split(text, "") {

		shift := strings.Index(alphabet, keyChars[count%keyLenght])
		index := strings.Index(alphabet, letter)

		if mode == "encrypt" {
			result += alphabetChars[(index+shift)%len(alphabet)]
		} else if mode == "decrypt" {
			result += alphabetChars[(index-shift)%len(alphabet)]
		} else {
			return result
		}

		count += 1
	}

	return result
}
