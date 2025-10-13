package main

import "fmt"
func main (){
	var x int
	var r, t, volume float32
	const pi = 3.14
	
	fmt.Print("Masukan berapa jml kerucut yang akan di hitung : ")
	fmt.Scan(&x)

	for i := 1; i<=x; i++ {
		fmt.Print("Masukan jari-jari alas kerucut : ")
		fmt.Scan(&r)
		fmt.Print("Masukan tinggi kerucut : ")
		fmt.Scan(&t)

		volume = 1/3.0 * pi * r * r * t
		fmt.Printf("Volume kerucut ke-%d adalah : %.14f\n", i, volume)
	}	
}