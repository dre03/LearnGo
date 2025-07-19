package main

import "fmt"

func main() {
	// Operasi matematika
	// comntohnya
	var a = 10
	var b = 10
	var c = a + b

	var kali = a * b

	fmt.Println(c)
	fmt.Println(kali)

	fmt.Println("===================")
	
	// operasi augmented assigments
	var i = 20
	i += 10 // sama aja kaya => i = i + 10
	fmt.Println(i)
	i -= 2
	fmt.Println(i)
	
	fmt.Println("===================")
	// unary operator atau untuk mempersingkat operator
	i++  
	fmt.Println(i)
	i--  
	fmt.Println(i)
	
	
}