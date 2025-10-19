package main 
import (
	"fmt"
)
func main() {
	for {
	var a, b int
	fmt.Print("masukan angka dan pangkatnya: ")
	fmt.Scan(&a, &b)

	if  a == 0 && b == 0 {
		fmt.Print("selesai.")
		break 
	}
	
	hasil := 1
	for i := 1; i <= b; i++ {
		hasil = hasil * a
	}
	fmt.Printf("%d pangkat %d = %d\n\n", a, b, hasil)
  }
}