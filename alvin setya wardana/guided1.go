package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("masukan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukan nilai b : ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
}
