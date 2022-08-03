package main

import "fmt"

func main() {
	var tinggi = map[string]int{"Aldo": 182, "Yosep": 178}
	tampil_mahasiswa(tinggi)
}

func tampil_mahasiswa(tinggi map[string]int) {
	fmt.Printf("Aldo : %d cm", tinggi["Aldo"])
	fmt.Printf("\nYosep : %d cm", tinggi["Yosep"])
}
