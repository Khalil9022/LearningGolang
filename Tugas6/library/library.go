package library

import "fmt"

type mahasiswa struct {
	nama string
	umur int
}

func Hasil() {
	var x mahasiswa
	x.nama = "Khalil"
	x.umur = 12
	x.Namasaya()
}

func (s mahasiswa) Namasaya() {
	fmt.Println(s.nama, s.umur)
}
