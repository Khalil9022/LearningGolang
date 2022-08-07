package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql-master"
	"net/http"
)

type mahasiswa struct {
	id       int
	Nama     string
	Jurusan  string
	Angkatan string
}

var data []mahasiswa

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/niomic_golang_tugas16")

	if err != nil {
		return nil, err
	}
	return db, nil
}

func ambil_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from tbl_mhs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = mahasiswa{}
		var err = rows.Scan(&each.id, &each.Nama, &each.Jurusan, &each.Angkatan)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func tampil_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	ambil_data()
	http.HandleFunc("/menu", tampil_data)
	fmt.Println("Menajalankan web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
