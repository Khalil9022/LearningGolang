package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	defer pulihkan_saya()
	fmt.Print("Masukkan Angka : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "ini adalah angka")
	} else {
		fmt.Println(input, " Bukan Angka ini")
		panic(err.Error())
		fmt.Println("Munculkan Saya")
	}
}

func pulihkan_saya() {
	if r := recover(); r != nil {
		fmt.Println("akhirnya saya ditampilkan :(\n", r)
	} else {
		fmt.Println("ini lancar saja")
	}
}
