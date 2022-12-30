package main

import (
	"fmt"
)

func main() {
	var message = "CSOITEUIWUIZNSROCNKFD"
	var key = "GOLANG"
	var keyLenth = len(key)
	var m = 0
	for l, _ := range message {
		i := rune(message[l])
		if m < keyLenth {
			k := rune(key[m])
			j := k - 65 + i - 65
			m++
			if j > 26 {
				j = j - 26 + 65
			} else {
				j = j + 65
			}
			fmt.Print(string(j))
		} else {
			m = m - keyLenth
		}
	}
}
