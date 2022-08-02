package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")
	fmt.Println("Jumlah Element : ", len(buah))
	fmt.Println("Isi element : ", buah)

	for i := 0; i < len(buah); i++ {
		fmt.Println("Element ke - : ", i, " ", buah[i])
	}
}
