package main

import "fmt"

func main() {
	//konversi tipe data pada go

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	
	fmt.Println("=============")



	// konversi string
	var name string = "Andre"
	var n byte = name[1]
	var nString = string(n)

	fmt.Println(name)
	fmt.Println(nString)
}