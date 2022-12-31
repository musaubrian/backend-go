package utils

import "math/rand"

var collection = []rune("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
func GenerateUrl()  string{
    str := make([]rune, 4)
    
    for i := range str {
        str[i] = collection[rand.Intn(len(collection))]
    }

    return string(str)
}
