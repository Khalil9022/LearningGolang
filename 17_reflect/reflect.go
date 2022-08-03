package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 10
	var reflectnumber = reflect.ValueOf(number)

	fmt.Println("Tipe Variable : ", reflectnumber.Type())

	if reflectnumber.Kind() == reflect.Int {
		fmt.Println("Nilai Variable : ", reflectnumber.Int())
	}
}
