package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "jeruk pisang ANGGUR2"
	var regex, err = regexp.Compile(`[a-z]+`)
	var regex2, err2 = regexp.Compile(`[A-Z]+\d`) // huruf A - Z dan harus ada nominal angka di salah satu katanya

	if err != nil {
		fmt.Println(err.Error())
	}
	var hasil = regex.FindAllString(text, 2)   //mengambil 2 kata saja
	var hasil2 = regex.FindAllString(text, -1) //mengambil semua kata
	fmt.Println(hasil)
	fmt.Println(hasil2)

	if err2 != nil {
		fmt.Println(err2.Error())
	}
	var hasil3 = regex2.FindAllString(text, -1)
	var hasil4 = regex2.MatchString(text)                // cek boolean
	var hasil5 = regex2.FindStringIndex(text)            // cek index pada kata
	var hasil6 = regex2.ReplaceAllString(text, "DURIAN") // replace string jadi kata baru
	fmt.Println(hasil3)
	fmt.Println(hasil4)
	fmt.Println(hasil5)
	fmt.Println(hasil6)
}
