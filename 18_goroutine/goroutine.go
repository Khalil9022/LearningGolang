package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go tampilkan_pesan(5, "saya kirim")
	tampilkan_pesan(5, "saya kedua")

	//blocking data, agar data dapat ditampilkan menggunakan go
	var input string
	fmt.Scanln(&input)
}

func tampilkan_pesan(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
