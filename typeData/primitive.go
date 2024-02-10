package main

import "fmt"

func main() {

}

func Primitive() {

	// Tipe data numerik && cara penulisan variabel 1
	var bilanganBulat int = 10         // Bilangan bulat
	var bilanganDesimal float64 = 3.14 // Bilangan desimal

	// Tipe data string && cara penulisan variabel 2
	nama := string("Golang") // String literal
	// nama = "javascript"

	// Tipe data boolean
	benar := bool(true) // Nilai boolean
	// benar = false

	// Deklarasi variabel tanpa `var`
	umur := 20 // Tipe data diinferensi dari nilai

	// Mencetak nilai variabel
	fmt.Println("Bilangan bulat:", bilanganBulat)
	fmt.Println("Bilangan desimal:", bilanganDesimal)
	fmt.Println("Nama:", nama)
	fmt.Println("Benar:", benar)
	fmt.Println("Umur:", umur)
}
