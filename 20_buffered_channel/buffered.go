package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	pesan := make(chan int, 3)

	go func() {
		for {
			i := <-pesan
			fmt.Println("terima data ", i)
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Println("Kirim data ", i)
		pesan <- i
	}
}
