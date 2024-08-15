package main

import (
	"encoding/csv" // Mengimpor paket csv untuk membaca dan menulis data CSV
	"os"           // Mengimpor paket os untuk operasi sistem seperti menulis ke stdout
)

func main() {
	// Membuat penulis CSV yang menulis ke stdout (konsol)
	writer := csv.NewWriter(os.Stdout)

	// Menulis baris pertama ke CSV
	_ = writer.Write([]string{"luffy", "d", "monkey"})
	// Menulis baris kedua ke CSV
	_ = writer.Write([]string{"roro", "noa", "zoro"})
	// Menulis baris ketiga ke CSV
	_ = writer.Write([]string{"vins", "moke", "sanji"})

	// Memastikan semua data ditulis ke output
	writer.Flush()
}
