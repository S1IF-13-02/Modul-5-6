package main

import (
	"fmt"
)

func main() {
	var j, n int
	var volume float64
	fmt.Print ("Masukkan bilangan: ")
	fmt.Scan(&n)
	
	for j = 1; j <= n ; j++{
		var jari_jari, tinggi int
		fmt.Scan(&jari_jari, &tinggi)

		volume = 1.0 / 3.0 *3.14 *float64(jari_jari)* float64(jari_jari)* float64(tinggi)
		fmt.Println(volume)
	}
	
}

