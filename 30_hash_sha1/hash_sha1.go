package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	//SHA1 => secure hash algoritma 1
	var text = "text ini rahasiax"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var enkripsi = sha.Sum(nil)
	var stringEnkripsi = fmt.Sprintf("%x", enkripsi)

	fmt.Println(stringEnkripsi)

}
