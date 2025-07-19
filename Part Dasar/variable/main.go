package main

import "fmt"

func main() {
	// tipe data string
	var name string
	name = "Andre Apriyana"
	fmt.Println(name)

	name = "Apri"
	fmt.Println(name)

	// tanpa menginisilisai type datanya
	var age = 22
	fmt.Println("umur" , name , "adalah" , age )

	// membuat variabel tanpa kata kunci var bisa dengan cara
	gender := "famale"
	gender = "male"
	fmt.Println(gender)

	// Membuat multiple variable
	var (
		city = "Bogor"
		country = "Indonesia"
	)
	fmt.Println(city, country)


}