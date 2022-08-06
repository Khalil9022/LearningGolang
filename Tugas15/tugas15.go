package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "perulangan dari angka 1 - 100 : ")
	for i := 1; i < 101; i++ {
		fmt.Fprint(w, "[", i, "] ")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang di website menggunakan GOLANG ini :D")
	})
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
