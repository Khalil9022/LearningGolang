package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var str = "156"
	var num, err = strconv.Atoi(str) // konversi string ke int
	var str2 = strconv.Itoa(num)     // konversi int to string
	if err == nil {
		fmt.Println(num)
		fmt.Println(str + "10")
	}

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(str2))
}
