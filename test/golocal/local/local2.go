package local

import "fmt"

func GoLocalD()  {
	go_local a int
	a++
	fmt.Println("goLocal:", a)
}
