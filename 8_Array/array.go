package main

import "fmt"

func main() {
	var buah = [4]string{"Apel", "Jeruk", "Anggur", "Nanas"}
	buah[1] = "Pisang"
	fmt.Println("Jumlah element : ", len(buah))
	fmt.Println("Isi element : ", buah)
}
