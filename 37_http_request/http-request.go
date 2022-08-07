//JALANKAN FILE PADA FOLDER 36_API/api.go terlebih dahulu untuk mengaktifkan localhostnya

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type menu_makanan struct {
	ID       string
	Namamenu string
	Harga    int
}

var baseURL = "http://localhost:8080"

func ambil_api() ([]menu_makanan, error) {
	var err error
	var client = &http.Client{}
	var data []menu_makanan

	requst, err := http.NewRequest("POST", baseURL+"/menu", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(requst)
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

func ambil_api_spesifik(menu string) (menu_makanan, error) {
	var err error
	var client = &http.Client{}
	var data menu_makanan

	var param = url.Values{}
	param.Set("Namamenu", menu)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/carimenu", payload)

	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
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
	fmt.Print("Masukkan menu yg ingin diliat : ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		var menu, err = ambil_api()
		if err != nil {
			fmt.Println("Menu tidak tersedia!", err.Error())
			return
		}

		for _, each := range menu {
			fmt.Println("ID : ", each.ID, "Menu : ", each.Namamenu, "Harga : ", each.Harga)
		}
	case 2:
		var menu2, err2 = ambil_api_spesifik("Lotek")
		if err2 != nil {
			fmt.Println("Menu Tidak tersedia")
			return
		}
		fmt.Println("ID : ", menu2.ID, "Menu : ", menu2.Namamenu, "Harga : ", menu2.Harga)
	}

}
