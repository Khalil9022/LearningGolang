package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var hasil, _ = exec.Command("ipconfig").Output()
	fmt.Println(string(hasil))
}
