package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan bilangan bulat: ")
	fmt.Scan(&n)

	var hasil int = 1 
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}

	fmt.Println("hasil faktorial dari", n, "adalah", hasil)
}
