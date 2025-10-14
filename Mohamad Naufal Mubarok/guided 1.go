package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("masukan bilangan a:")
	fmt.Scan(&a)

	fmt.Print("masukan bilangan b:")
	fmt.Scan(&b)

    for i := a; i <= b; i++{
	fmt.Println("keluaran", i)
  }
}