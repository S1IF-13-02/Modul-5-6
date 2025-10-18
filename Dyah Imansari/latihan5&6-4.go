package main
import "fmt"

func main() {
	var b, i int
	var faktorial uint64
	faktorial = 1
	
	fmt.Print("Masukkan bilangan bulat non-negatif : ")
	fmt.Scanln(&b)
	for i = 1; i <= b; i++ {
		faktorial = faktorial * uint64(i)
	}
	fmt.Println("Hasil faktorial :", faktorial)
}