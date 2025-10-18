package main
import "fmt"

func main() {
	var basis, pangkat, i int
	var hasil int
	hasil = 1

	fmt.Print("Masukkan basis: ")
	fmt.Scanln(&basis)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&pangkat)
	for i = 0; i < pangkat; i++ {
		hasil = hasil * basis
	}
	fmt.Println("Hasil pangkat =", hasil)
}