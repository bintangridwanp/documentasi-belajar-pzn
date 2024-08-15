package main

import (
	"encoding/csv" // Mengimpor paket csv untuk membaca dan menulis data CSV
	"fmt"          // Mengimpor paket fmt untuk mencetak output ke konsol
	"io"           // Mengimpor paket io untuk operasi input/output
	"strings"      // Mengimpor paket strings untuk memanipulasi string
)

func main() {
	// Mendefinisikan string CSV
	csvString := "luffy d monkey\n" +
		"roronoa zoro\n" +
		"vismoke sanji"

	// Membuat pembaca CSV dari string
	reader := csv.NewReader(strings.NewReader(csvString))

	// Membaca setiap baris CSV dalam loop
	for {
		record, err := reader.Read()
		if err == io.EOF {
			// Jika mencapai akhir file, keluar dari loop
			break
		}

		// Mencetak setiap record (baris) yang dibaca
		fmt.Println(record)
	}
}
