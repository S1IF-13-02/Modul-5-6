package main

import "fmt"

func main() {
	var n, hasil int


	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var bilangan int

		fmt.Printf("Masukkan bilangan ke-%d: ", i)
		fmt.Scan(&bilangan)

		for j :=1; j <= bilangan; j++{
			hasil = hasil + 1
		}
	}
	
	fmt.Print("Hasil: ", hasil)
}
