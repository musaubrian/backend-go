// Package utils holds the GenerateUrl function
package utils

import "crypto/rand"

var collection = []byte(
	"0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

// Returns a 4 characters string
// This will be used to redirect to the actual url
func GenerateUrl() string {

	linkLength := len(collection)
	shortUrl := make([]byte, 4)
	rand.Read(shortUrl)

	for i := 0; i < 4; i++ {
		shortUrl[i] = collection[int(shortUrl[i]%byte(linkLength))]
	}
	return string(shortUrl)
}
