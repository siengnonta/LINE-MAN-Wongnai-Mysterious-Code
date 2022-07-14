package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	var whatIsIt string
	for i := len(sd) - 1; i >= 0; i-- {
		character := rune(sd[i])
		whatIsIt += string(character)
	}
	
	fmt.Println(whatIsIt) // Join:us:at:LINE:MAN:Wongnai
}
