package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hai") // defer digunakan untuk MENG AKHIRKAN baris sebuah kode, jadi ini akan tampil di akhir
	fmt.Println("selamat datang")
	defer fmt.Println("YESSIRR")
	fmt.Println("Dirumah saya")
	tampilkan()
	os.Exit(1) // untuk mengehntikan program yg sedang berjalan. jika menggunakan Exit, maka defer juga tidak akan bisa jalan
}

func tampilkan() {
	fmt.Println("Saya ditampilkan")
}
