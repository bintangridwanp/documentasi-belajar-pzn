package main

import (
	"fmt"  // Mengimpor paket fmt untuk mencetak output ke konsol
	"path" // Mengimpor paket path untuk operasi pada path file
)

func main() {
	// Mencetak direktori dari path "hello/world.go"
	fmt.Println(path.Dir("hello/world.go")) // Output: "hello"

	// Mencetak nama file dari path "hello/world.go"
	fmt.Println(path.Base("hello/world.go")) // Output: "world.go"

	// Mencetak ekstensi file dari path "hello/world.go"
	fmt.Println(path.Ext("hello/world.go")) // Output: ".go"

	// Menggabungkan beberapa elemen path menjadi satu path
	fmt.Println(path.Join("hello", "world", "main.go")) // Output: "hello/world/main.go"
}
