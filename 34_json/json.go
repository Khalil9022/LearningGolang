package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	//json harus pakai huruf besar diawalan kata pada golang
	Nama  string `json:"Nama"`
	Umur  int
	Kelas string `json:"Kelas"`
}

func main() {
	var jsonString = `{"Nama" : "Ilham", 
					   "Umur" : 18,
					   "Kelas" : "XII"}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nama :", data.Nama)
	fmt.Println("Umur :", data.Umur)
	fmt.Println("Kelas :", data.Kelas)

	decodJson()
}

func decodJson() {
	var object = []User{{"Ilham", 18, "XII"}, {"Abdul", 19, "XI"}}
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
