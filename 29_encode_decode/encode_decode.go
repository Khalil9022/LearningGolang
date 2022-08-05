package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var nama_saya = "Ilham"
	var encodeString = base64.StdEncoding.EncodeToString([]byte(nama_saya))
	fmt.Println("encodeing stringnya :", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println(decodeByte)
	fmt.Println("isi dari stringnya adalah :", decodeString)
}
