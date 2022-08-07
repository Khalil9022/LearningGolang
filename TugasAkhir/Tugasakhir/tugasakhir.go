//JALANKAN FILE PADA tugas 17 terlebih dahulu untuk mengaktifkan API nya

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type mahasiswa struct {
	Nama     string
	Jurusan  string
	Angkatan string
}

var baseURL = "http://localhost:8080"

func getmahasiswa() ([]mahasiswa, error) {
	var err error
	var client = &http.Client{}
	var data []mahasiswa

	request, err := http.NewRequest("POST", baseURL+"/mhs", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getspesifikmahasiswa(nama string) (mahasiswa, error) {
	var err error
	var client = &http.Client{}
	var data mahasiswa

	var param = url.Values{}
	param.Set("Nama", nama)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/carimhs", payload)

	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	respone, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer respone.Body.Close()

	err = json.NewDecoder(respone.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func deletemhs(nama string) (mahasiswa, error) {
	var err error
	var client = &http.Client{}
	var data mahasiswa

	var param = url.Values{}
	param.Set("Nama", nama)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/deletemhs", payload)

	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	respone, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer respone.Body.Close()

	err = json.NewDecoder(respone.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func main() {
	var input int
	fmt.Println("Menu : ")
	fmt.Println("1. Cari semua")
	fmt.Println("2. Cari spesifik")
	fmt.Println("3. Update Data")
	fmt.Println("4. Delete Data")
	fmt.Print("Masukkan menu yg ingin diliat : ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		var mhs, err = getmahasiswa()
		if err != nil {
			fmt.Println("Tidak ditemukan mahasiswa!", err.Error())
			return
		}

		for _, each := range mhs {
			fmt.Println("Nama : ", each.Nama, "Jurusan : ", each.Jurusan, "Angkatan : ", each.Angkatan)
		}

	case 2:
		var mhs, err2 = getspesifikmahasiswa("Gigi")
		if err2 != nil {
			fmt.Println("Mahasiswa Tidak tersedia")
			return
		}
		fmt.Println("Nama : ", mhs.Nama, "Jurusan : ", mhs.Jurusan, "Angkatan : ", mhs.Angkatan)

	case 3:

	case 4:
		var mhs, err3 = deletemhs("Akino")
		if err3 != nil {
			fmt.Println("Mahasiswa Berhasil dihapus ", mhs)
			return
		}
	}
}
