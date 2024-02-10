package main

import "fmt"

func NonPrimitive() {
	// Array | jumlah panjang array & tipe data sudah ditentukan sejak awal dan tidak boleh berubah
	var fruits [3]string // Array dengan 3 elemen bertipe string
	fruits[0] = "Apel"
	fruits[1] = "Pisang"
	fruits[2] = "Jeruk" // jika tidak diisi maka akan menjadi nil

	// Slice
	var numbers []int                  // Slice integer yang dapat diubah ukuran panjang nya
	numbers = append(numbers, 1, 2, 3) // Menambahkan elemen ke slice

	// NOTE: baik Array maupun SLice keduanya diakses dan di assign menggunakan index

	// Map
	// var person map[string]string // Map yang memetakan string ke string |
	// person["nama"] = "Andi"
	// person["umur"] = "25"
	// person["test"] = "test"

	// NOTE: sedangkan Map diakses dan diassign menggunakan (Key value) pair mirip seperti di javascript tapi menggunakan [] seperti array

	// Mencetak nilai
	fmt.Println("Fruits:", fruits)
	fmt.Println("Numbers:", numbers[0])
	// fmt.Println("Person:", person)

}
