package library

import "fmt"

func iniprivate() {
	fmt.Println("ini private")
}

func Inipublic() {
	fmt.Println("ini public")
	iniprivate()
}
