package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	var tipe1 = "Khalil"
	var tipe2 = 22
	runtime.GOMAXPROCS(2)
	//reflect
	bacadata(tipe1, tipe2)

	//goroutine
	go bacadata(tipe1, tipe2)

	var input string
	fmt.Scanln(&input)
}

func bacadata(kata string, number int) {
	var reflectnumber = reflect.ValueOf(number)
	var reflectkata = reflect.ValueOf(kata)

	fmt.Println(reflectnumber.Type())
	fmt.Println(reflectkata.Type())
}
