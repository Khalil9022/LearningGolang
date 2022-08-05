package main

import (
	"fmt"
	"regexp"
)

func main() {
	kata := "Hello world"
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hasil = regex.FindAllString(kata, 1)
	fmt.Println(hasil)
	var hasil2 = regex.MatchString(kata)
	fmt.Println(hasil2)

	var input string
	fmt.Print("Masukkan kata yg ingin dimasukkan : ")
	fmt.Scanln(&input)

	var hasil3 = regex.FindAllString(input, -1)
	fmt.Println(hasil3)

}
