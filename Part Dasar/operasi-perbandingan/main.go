package main

import "fmt"

func main() {

	// Perasi perbandinggan di go
	var name1 = "Andre"
	var name2 = "Andre"

	var result bool = name1 == name2

	fmt.Println("Perbandingan 2 nilai tipe data string")
	fmt.Println(result)
	
	fmt.Println("================")
	fmt.Println("Perbandingan 2 nilai tipe data int")
	var nilai1 = 100
	var nilai2 = 200

	fmt.Println(nilai1 > nilai2)
	fmt.Println(nilai1 < nilai2)
	fmt.Println(nilai1 == nilai2)
	fmt.Println(nilai1 != nilai2)
	fmt.Println("================")

	var value = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Println("nilai", value, isEqual)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}