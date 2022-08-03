package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Angka ke -", i)
	}

	//Cara ke 2 :
	var j = 0
	for j < 5 {
		fmt.Println("Angka J ke-", j)
		j++
	}
}
