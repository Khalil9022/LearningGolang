package main

import (
	"fmt"
	"strings"
)

func main() {
	var apakahada = strings.Contains("disini adalah rumah", "mah")      // menampilkan true atau false jika kata "mah" ada di dalam kata
	var apakahada2 = strings.HasPrefix("disini adalah rumah", "disini") //melihat awalan kata saja. "disini" saja. d - di - dis - disi dll
	var apakahada3 = strings.HasSuffix("disini adalah rumah", "mah")    // melihat akhiran kata saja. "Rumah". h - ah- mah- umah- rumah
	var berapabanyak = strings.Count("misalnya ini tulisan", "t")       // mencari berapa banyak huruf t
	var index1 = strings.Index("misalnya ini tulisan", "i")             //mencari index i keberapa. awalnya saja
	fmt.Println(apakahada)
	fmt.Println(apakahada2)
	fmt.Println(apakahada3)
	fmt.Println(berapabanyak)
	fmt.Println(index1)
	cari("buah apel", "apel") // untuk replace kata kata pada substring yg kita inputkan
}

func cari(text string, cari string) {
	var textbaru = strings.Replace(text, cari, "nanas", 1)
	fmt.Println(textbaru)
}
