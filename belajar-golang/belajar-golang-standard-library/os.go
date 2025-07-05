package main

import (
	"fmt"
	"os"
)

func main() {

	// Menggunakan os.Args untuk mengambil argumen dari baris perintah
	args := os.Args

	// os.Args adalah slice yang berisi argumen baris perintah
	for _, arg := range args {
		fmt.Println(arg)
	}

	// Menggunakan os.Getenv untuk mengambil nilai dari variabel lingkungan
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Sedang ERROR", err.Error())
	}

}
