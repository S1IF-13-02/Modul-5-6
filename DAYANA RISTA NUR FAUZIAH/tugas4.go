package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan bilangan bulat non-negatif (n): ")
	fmt.Scan(&n)

	faktorial := 1

	for i := 1; i <= n; i++ {

		faktorial *= i
	}

	fmt.Println(faktorial)
}
