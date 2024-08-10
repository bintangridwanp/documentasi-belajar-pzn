package main

import (
	"fmt"
	"strings"
)

func main() {
	// Memeriksa apakah string "Luffy D Monkey" mengandung kata "Luffy"
	fmt.Println(strings.Contains("Luffy D Monkey", "Luffy"))

	// Memisahkan string "Luffy D Monkey" menjadi slice berdasarkan spasi
	fmt.Println(strings.Split("Luffy D Monkey", " "))

	// Mengubah string "Luffy D Monkey" menjadi huruf kecil
	fmt.Println(strings.ToLower("Luffy D Monkey"))

	// Mengubah string "Luffy D Monkey" menjadi huruf besar
	fmt.Println(strings.ToUpper("Luffy D Monkey"))

	// Menghilangkan spasi di awal dan akhir dari string "      Luffy D Monkey      "
	fmt.Println(strings.Trim("      Luffy D Monkey      ", " "))

	// Mengganti semua kata "Luffy" dengan "Budi" dalam string "Luffy D Monkey Luffy Khannedy"
	fmt.Println(strings.ReplaceAll("Luffy D Monkey Luffy Khannedy", "Luffy", "Budi"))
}
