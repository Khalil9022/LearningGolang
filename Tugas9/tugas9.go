package main

import "fmt"

func main() {
	var angka int
	var err error
	defer pulihkan()

	fmt.Print("Masukkan angka POSITIF : ")
	fmt.Scanln(&angka)

	if angka >= 0 {
		fmt.Println("Angka ", angka, " adalah positif")
	} else {
		fmt.Println("Angka ", angka, " tidak positif bro")
		panic(err.Error())
	}
}

func pulihkan() {
	if r := recover(); r != nil {
		fmt.Println("SAYA DITAMPILKAN BROOK \n", r)
	} else {
		fmt.Println("Angkanya benar")
	}
}
