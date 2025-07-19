package main

import "fmt"

func main() {

	// Aturan membuat konstanta itu, diawali degngan kata kunci const den diikuti dengan namanya
	// konstanta sekali di deklarisikan itu datanya tidak bisa diubah dan datanya bakalan error

	var name = "Andre Apri"
	const nim = 011021111

	// dibawah ini error karena datanya tidak bisa diubah
	// nim = 923643

	fmt.Println("Nama" , name ,  "nim" , nim)

	// 
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.2f\n", decimalNumber)
}

