package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql-master"
	"net/http"
)

type menu_makanan struct {
	ID       string
	Namamenu string
	Harga    int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/golang_api_menumakanan")

	if err != nil {
		return nil, err
	}
	return db, nil
}

// var data = []menu_makanan{
// 	{"C01", "Karedok", 20000},
// 	{"C03", "Pecel Lele", 15000},
// 	{"C04", "Ketoprak", 17000},
// 	{"C09", "Ayam Bakar", 25000},
// }

var data []menu_makanan

func ambil_menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

func cari_menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var namamenu = r.FormValue("Namamenu")
		var result []byte
		var err error

		for _, each := range data {
			if each.Namamenu == namamenu {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}

		http.Error(w, "Menu makanan tidak tersedia", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func ambil_data() {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from tbl_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu_makanan{}
		var err = rows.Scan(&each.ID, &each.Namamenu, &each.Harga)

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

func main() {
	ambil_data()
	http.HandleFunc("/menu", ambil_menu)
	http.HandleFunc("/carimenu", cari_menu)
	fmt.Println("Menajalankan web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
