package main

import (
	"fmt"
	"runtime"
)

var pesan = make(chan string) //chan = channel

func main() {
	runtime.GOMAXPROCS(2) //membuat core processor

	//3 buah goroutine jadi 1 chanel
	go hello("a")
	go hello("b")
	go hello("c")

	var message1 = <-pesan
	fmt.Println(message1)

	var message2 = <-pesan
	fmt.Println(message2)

	var message3 = <-pesan
	fmt.Println(message3)
}

func hello(nama string) {
	var data = fmt.Sprintf("Halo %s", nama)
	pesan <- data
}
