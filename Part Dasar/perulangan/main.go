package main

import (
	"fmt"
)

func main() {

	// Perulangan Menggunakan Keyword for
	// PerulanganFor()
	
	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	// PerulanganDenganArgumen()
	
	// Penggunaan Keyword for Tanpa Argumen => dengan break
	// forTanpaArgumen()
	
	// Penggunaan Keyword for - range
	// ForRangeString()
	// ForArray()
	// ForRSlice()
	// ForRMap()
	// NumerikRange()

	// Penggunaan Keyword break & continue
	// ForBreakContinue()

	// Perulangan Bersarang
	// ForBersarang()
	// Perulanagan Label
	ForLabel()

}

// Perulangan Menggunakan Keyword for
func PerulanganFor()  {
	fmt.Println("Perulangan Cara 1")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("====================")
}

// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
func PerulanganDenganArgumen()  {
	fmt.Println("Perulangan Cara 2")
	var angka = 1
	for angka < 5{
		fmt.Println("Angka",  angka)
		angka++
	}
	fmt.Println("====================")
}

// Penggunaan Keyword for Tanpa Argumen => dengan break
func forTanpaArgumen()  {
	fmt.Println("Perulangan Cara 3")
	angka := 0
	for {
		fmt.Println("Angka", angka)
		angka++
		if angka == 4 {
			break
		}
	}
	fmt.Println("====================")
}

// Penggunaan Keyword for - range
func ForRangeString()  {
	fmt.Println("Perulangan Cara 4")
	var xs = "123"
	for i, v := range xs{
		fmt.Println("Index=", i , "Value=", v,)
	}
	fmt.Println("====================")
	
	
}

func ForArray()  {
	var arr = []int{10,20,30,40,50} // array
	for in, val := range arr{
		fmt.Println("Index", in, "Valeu", val)
	}
	fmt.Println("====================")
}

func ForRSlice()  {
	var arr = [5]int{10, 20, 30, 40, 50} // array
	var zs = arr[3:4]
	for i, val := range zs{
		fmt.Println(i, "Valeu", val)
	}
}

func ForRMap()  {
	var obj = map[byte]int{'a':0, 'b':1, 'c':2, 'd':3} //map
	for key, val := range obj{
		el := string(key)
		fmt.Println("Key", el, "Valeu", val)
	}
	fmt.Println("===========")
}

// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
func NumerikRange()  {
	for i:= range 5{
		fmt.Println(i)
	}
}

// Penggunaan Keyword break & continue
func ForBreakContinue()  {
	// Menentukan atau mencari bilanagan genap
	for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}

// Perulangan Bersarang
func ForBersarang()  {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j," ")
		}
		fmt.Println()
	}

}

// Pemanfaatan Label Dalam Perulangan
func ForLabel()  {
	label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break label
			}
			fmt.Print("matriks [",i,"][",j,"b]\n")
		}
	}	
}
	