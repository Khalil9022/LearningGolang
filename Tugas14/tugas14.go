package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var input string
	var hasil1, _ = exec.Command("help").Output()
	var hasil2, _ = exec.Command("ipconfig").Output()
	var hasil3, _ = exec.Command("getmac").Output()

	fmt.Println("Select wich one you prefer to look : ")
	fmt.Println("1. Hasil 1 (help) ")
	fmt.Println("2. Hasil 2 (ipconfig) ")
	fmt.Println("3. Hasil 3 (getmac) ")
	fmt.Print("Input your number : ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println(string(hasil1))
	case "2":
		fmt.Println(string(hasil2))
	case "3":
		fmt.Println(string(hasil3))
	default:
		fmt.Print("Wrong answer")
	}
}
