package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//b := make([]byte, 32)
	//_, err := rand.Read(b)
	//if err != nil {
	//	return
	//}

	bytes, err := bcrypt.GenerateFromPassword([]byte("123321"), 14)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
