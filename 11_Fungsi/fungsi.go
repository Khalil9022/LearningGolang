package main

import "fmt"

func main() {
	var x1, x2 = tampilkan_multi_fungsi(10, 5)
	fmt.Println(tampilkan_pesan())
	fmt.Println(penjumlahan(10, 6))
	println(x1)
	println(x2)
}

func tampilkan_pesan() string {
	fmt.Println("pesan berhasil diterima :D")
	return "pesan bisa ditampilkan jika diterurn ya bro :D"
}

func penjumlahan(x int, y int) int {
	return x + y
}

func tampilkan_multi_fungsi(x int, y int) (int, int) {
	var z = x * y
	var a = x / y
	return z, a
}
