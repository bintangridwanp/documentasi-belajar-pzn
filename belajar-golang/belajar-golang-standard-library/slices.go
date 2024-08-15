package main

import (
	"fmt"    // Mengimpor paket fmt untuk mencetak output ke konsol
	"slices" // Mengimpor paket slices untuk operasi pada slice
)

func main() {
	// Mendefinisikan slice string
	names := []string{"zoro", "sanji", "nami", "luffi"}
	// Mendefinisikan slice integer
	values := []int{100, 95, 80, 87}

	// Mencetak elemen terkecil dalam slice names
	fmt.Println(slices.Min(names)) // Output: "luffi"

	// Mencetak elemen terkecil dalam slice values
	fmt.Println(slices.Min(values)) // Output: 80

	// Mencetak elemen terbesar dalam slice names
	fmt.Println(slices.Max(names)) // Output: "zoro"

	// Mencetak elemen terbesar dalam slice values
	fmt.Println(slices.Max(values)) // Output: 100

	// Mengecek apakah slice names mengandung elemen "nami"
	fmt.Println(slices.Contains(names, "nami")) // Output: true

	// Mencari indeks elemen "sanji" dalam slice names
	fmt.Println(slices.Index(names, "sanji")) // Output: 1

	// Mencari indeks elemen "zoro" dalam slice names
	fmt.Println(slices.Index(names, "zoro")) // Output: 0
}
