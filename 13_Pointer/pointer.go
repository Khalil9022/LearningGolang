package main

import "fmt"

func main() {
	var nomor int = 10
	var alamat_nomor *int = &nomor
	var tas string = "Hello world"
	var alamat_tas *string = &tas

	fmt.Println("Nilai dari nomor : ", nomor)
	fmt.Println("Alamat Variabel nomor : ", alamat_nomor)

	fmt.Println("Nilai dari nomor : ", tas)
	fmt.Println("Alamat Variabel nomor : ", alamat_tas)
}
