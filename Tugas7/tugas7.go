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

	//goroutine
	go bacadata1(tipe1, "ke-1")
	go bacadata2(tipe2, "ke-1")

	//reflect
	bacadata1(tipe1, "ke-2")
	bacadata2(tipe2, "ke-2")

	var input string
	fmt.Scanln(&input)
}

func bacadata1(kata string, pesan string) {
	var reflectkata = reflect.ValueOf(kata)
	fmt.Println(pesan, reflectkata.Type())
}

func bacadata2(number int, pesan string) {
	var reflectnumber = reflect.ValueOf(number)
	fmt.Println(pesan, reflectnumber.Type())
}
