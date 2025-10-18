package main
import "fmt"

func main() {
	var n, j, jumlah int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)

	jumlah = 0
	for j = 1; j <= n; j++ {
		jumlah = jumlah + j
	}

	fmt.Println("Hasil penjumlahan =", jumlah)
}
