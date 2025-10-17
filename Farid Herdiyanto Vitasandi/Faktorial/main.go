package main

import "fmt"

func main() {
	
	var n int
	
	fmt.Print("Masukkkan bilangan: ")
	fmt.Scan(&n)

	faktorial := 1

	for i := 1; i <= n; i++{
		faktorial = faktorial * i
	}

	fmt.Printf("Hasil Faktorial dari bilangan %d adalah: %d", n, faktorial)
}
