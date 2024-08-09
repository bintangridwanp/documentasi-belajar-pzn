package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args // Mengambil argumen baris perintah yang diberikan saat menjalankan program

	for _, arg := range args { // Mengiterasi setiap argumen
		fmt.Println(arg) // Menampilkan argumen ke konsol
	}

	hostname, err := os.Hostname() // Mengambil nama host dari sistem
	if err == nil {
		fmt.Println(hostname) // Menampilkan nama host jika tidak ada error
	} else {
		fmt.Println("Sedang ERROR", err.Error()) // Menampilkan pesan error jika terjadi masalah
	}

}
