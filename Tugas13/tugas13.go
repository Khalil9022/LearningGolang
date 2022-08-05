package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama_kita = "Khalil Attalla"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(nama_kita))
	var sha = sha1.New()
	sha.Write([]byte(encodeString))
	var enkripsi = sha.Sum(nil)
	var stringEnkripsi = fmt.Sprintf("%x", enkripsi)

	fmt.Println("Kata yang akan di hash & encode :", nama_kita)
	fmt.Println("Hasil jadinya adalah :", stringEnkripsi)
}
