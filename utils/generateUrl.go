//Package utils deals with the general functions
package utils

import "crypto/rand"

//collection creates 'group' of characters and digits
var collection = []byte ("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

//GenerateUrl returns a 4 characters string
//This will be used to redirect to the actual url
func GenerateUrl()  string{

    linkLength := len(collection)
    str := make([]byte, 4)
    rand.Read(str)

    for i := 0; i < 4; i++{
        str[i] = collection[int(str[i]%byte(linkLength))]
    }
    return string(str)
}
