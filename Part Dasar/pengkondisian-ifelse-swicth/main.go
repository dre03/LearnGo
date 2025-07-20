package main

import "fmt"

func main() {
	// Belajar seleksi pengkondisian

	nilai := 8

	if nilai == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	}else if nilai > 5 {
		fmt.Println("lulus")
	}else if nilai == 4 {
		fmt.Println("hampir lulus")
	}else{
		fmt.Printf("tidak lulus. nilai anda %d\n", nilai)
	}

	fmt.Println("===================")
	// Variabel Temporary Pada if - else
	var point = 8840.0
	
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
		}else if percent >= 70 {
			fmt.Printf("%.1f%s good\n", percent, "%")
	}else{
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
	
	fmt.Println("===================")
	// Pengkondisian dengan Switch
	
	var hasil = 4

	switch hasil{
	case 8:
		fmt.Println("perfect")
	case 4,7,6,3,5:
		fmt.Println("awesome")
	// Kurung Kurawal Pada Keyword case & default
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
    	}
	}
	
	fmt.Println("===================")
	// Switch Dengan Gaya if - else
	nilai = 6
	switch{
	case nilai == 8:
		fmt.Println("perfect")
	case (nilai > 3 && nilai < 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
    	fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
    	}
	}

	// Seleksi Kondisi Bersarang dan bisa gabungan dari if dan switch
	fmt.Println("================")
	nilai = 0

	if nilai > 7 {
		switch nilai {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	}else{
		if nilai == 5{
			fmt.Println("nnot bad")
		}else if nilai == 3{
			fmt.Println("keep trying")
		}else{
			fmt.Println("you can do it")
			if nilai == 0 {
				fmt.Println("try herder!")
			}
		}
	}


	
}