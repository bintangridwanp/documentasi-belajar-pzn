package main

import (
	"fmt"    // Mengimpor paket fmt untuk mencetak output ke konsol
	"regexp" // Mengimpor paket regexp untuk bekerja dengan regular expressions
)

func main() {
	// Membuat variabel regex yang berisi pola regex `e([a-z])o`
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	// Mengecek apakah string "eko" cocok dengan pola regex
	fmt.Println(regex.MatchString("eko")) // Output: true

	// Mengecek apakah string "edo" cocok dengan pola regex
	fmt.Println(regex.MatchString("edo")) // Output: true

	// Mengecek apakah string "eKo" cocok dengan pola regex
	fmt.Println(regex.MatchString("eKo")) // Output: false (karena 'K' adalah huruf besar)

	// Mencari semua string dalam teks yang cocok dengan pola regex
	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
	// Output: ["eko", "edo", "ego", "eto"]
}
