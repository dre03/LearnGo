package main

import "fmt"

func main() {
	// Type Declarations adalah kemampuan membuat ulang tippe data baru dari tipe data yang sudah ada
	// type declaration biasanya digunakan untuk membuat alias terhadap tide data yanng sudah ada dengan tujuan agar lebih mudah dimengerti

	// contohnya
	type NoKTP string
	type Married bool

	var noKtpUser NoKTP = "31247892367345"
	var iSMarried Married = false

	fmt.Println(noKtpUser)
	fmt.Println(iSMarried)
	// Pada intinya kita bisa membuat tipe data baru dengan menngunakan tipe data yang sudah ada pada go
}