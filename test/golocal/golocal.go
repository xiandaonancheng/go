package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go callGoLocal(1)
	time.Sleep(time.Second)
	go callGoLocal(2)
	time.Sleep(time.Second)
	go callGoLocal(3)
	time.Sleep(time.Second)
}

func callGoLocal(callID int) {
	fmt.Println("call GoLocal:", callID)
	fmt.Println("---------------------go_local-------------------")
	goLocal()
	goLocal()
	goLocal()
	goLocal()
	goLocal()
	fmt.Println("---------------------runtime.NewGoLocal-------------------")
	goLocalA()
	goLocalB()
	goLocalA()
	goLocalB()
	goLocalA()
	goLocalB()
	fmt.Println("----------------------------------------")
}

func goLocal() {
	var a int
	a++
	fmt.Println("goLocal:", a)
}

func goLocalA() {
	a, _ := runtime.NewGoLocal[int]("go local", func() int {
		return 10
	})
	a.Val++
	fmt.Println("goLocalA:", a.Val)
}

func goLocalB() {
	b, _ := runtime.NewGoLocal[int]("go local", func() int {
		return 10
	})
	b.Val++
	fmt.Println("goLocalB:", b.Val)
}
