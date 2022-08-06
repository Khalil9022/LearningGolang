package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type pelajar struct {
	id    int
	nama  string
	umur  int
	kelas int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sql_tampil() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	//var umur = 18
	rows, err := db.Query("select * from tbl_belajar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pelajar
	//mengambil datanya menggunakan FOR :
	for rows.Next() {
		var each = pelajar{}
		var err = rows.Scan(&each.id, &each.nama, &each.umur, &each.kelas)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.nama, "Kelas :", each.kelas, "Umur :", each.umur)
	}
}

func sql_tambah() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tbl_belajar values (?,?,?,?)", nil, "Adi sudianto", "20", "12")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ubahdata() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update tbl_belajar set umur=? where nama =?", 20, "Susi Susanti")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("udah berhasil update datanya nih :D")
}

func deletedata() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("delete from tbl_belajar where nama=?", "Susi Susanti")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("udah berhasil Delete datanya nih :D")
}

func main() {
	var input int
	fmt.Println("Pilihan : ")
	fmt.Println("1. Tambah")
	fmt.Println("2. Update Data")
	fmt.Println("3. Delete Data")
	fmt.Println("4. Tampil Data")
	fmt.Print("Masukkan nomor yg ingin di pilih : ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		sql_tambah()
	case 2:
		ubahdata()
	case 3:
		deletedata()
	case 4:
		sql_tampil()
	default:
		fmt.Println("nomor yg diinputkan salah")
	}
}
