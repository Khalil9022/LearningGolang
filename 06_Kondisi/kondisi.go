package main

import "fmt"

func main() {
	var nilai = 5
	//Penggunaan IF ELSE
	if nilai == 10 {
		fmt.Println("Selamat, anda lulus brooo")
	} else if nilai > 7 {
		fmt.Println("Lumayan lah brooo")
	} else if nilai <= 7 {
		fmt.Println("Wah,,, harus bisa lebih tinggi ya brooo :D ")
	}

	//PEnggunaan Switch Case
	switch nilai {
	case 10:
		fmt.Println("Sangat GG broo")
	case 9:
		fmt.Println("Mantapp broo")
	case 8:
		fmt.Println("kurang brooo")
	default:
		fmt.Println("kurang bagus nilainya, tingkatkan lagi yaa :D")
	}
}
