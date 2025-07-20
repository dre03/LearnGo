package main

import "fmt"

func main() {

	// Operator boolean enghasilkan true atau false

	nilai := 80
	absensi := 83

	lulusUjian := nilai >= 80
	lulusAbsensi := absensi >= 77

	hasilAkhir := lulusUjian && lulusAbsensi

	fmt.Println(hasilAkhir)
	fmt.Println("=========")
	//  dengan cara singkat
	fmt.Println(nilai >= 80 && absensi >= 77)

}