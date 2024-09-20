package local

import "fmt"

func GoLocalA()  {
	go_local a int
	a++
	fmt.Println("goLocal:", a)
}

func GoLocalB()  {
	go_local b int
	b++
	fmt.Println("goLocal:", b)
}

func GoLocalC()  {
	go_local a int
	a++
	fmt.Println("goLocal:", a)
}
