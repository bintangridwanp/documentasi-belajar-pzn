package main

import (
	"encoding/base64" // Mengimpor paket base64 untuk encoding dan decoding base64
	"fmt"             // Mengimpor paket fmt untuk mencetak output ke konsol
)

func main() {
	// Mendefinisikan string yang akan di-encode
	value := "Luffy D Monkey"

	// Melakukan encoding string ke format base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded) // Output: "RWtvIEt1cm5pYXdhbiBLaGFubmVkeQ=="

	// Melakukan decoding string dari format base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		// Jika terjadi error saat decoding, cetak pesan error
		fmt.Println("Error", err.Error())
	} else {
		// Jika decoding berhasil, cetak hasil decoding
		fmt.Println(string(decoded)) // Output: "Eko Kurniawan Khannedy"
	}
}
