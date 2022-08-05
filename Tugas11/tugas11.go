package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "15"
	var str2 = "3"
	var num1, err1 = strconv.Atoi(str1)
	var num2, err2 = strconv.Atoi(str2)

	if err1 != nil || err2 != nil {
		fmt.Println(err1.Error())
	}
	var penjumlahan = num1 + num2
	var pengurangan = num1 - num2
	var pembagian = num1 / num2
	var perkalian = num1 * num2

	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(pembagian)
	fmt.Println(perkalian)
}
