package library

import "fmt"

type mahasiswa struct {
	nama string
	umur int
}

func Hasil(nama string, umur int) {
	var x mahasiswa
	x.nama = nama
	x.umur = umur
	x.Namasaya()
}

func (s mahasiswa) Namasaya() {
	fmt.Println(s.nama, s.umur)
}
