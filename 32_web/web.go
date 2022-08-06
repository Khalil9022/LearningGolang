package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar Teman!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " Hello")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Memulai webserver pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
