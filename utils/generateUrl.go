package utils

import "crypto/rand"

var collection = []byte ("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
func GenerateUrl()  string{

    linkLength := len(collection)
    str := make([]byte, 4)
    rand.Read(str)

    for i := 0; i < 4; i++{
        str[i] = collection[int(str[i]%byte(linkLength))]
    }
    return string(str)
}
