package main

import (
	"bufio" // Mengimpor paket bufio untuk buffered I/O
	"os"    // Mengimpor paket os untuk operasi sistem seperti menulis ke stdout
)

func main() {
	// Membuat penulis buffered yang menulis ke stdout (konsol)
	writer := bufio.NewWriter(os.Stdout)

	// Menulis string "hello world\n" ke buffer
	_, _ = writer.WriteString("hello world\n")

	// Menulis string "Selamat belajar\n" ke buffer
	_, _ = writer.WriteString("Selamat belajar\n")

	// Memastikan semua data dalam buffer ditulis ke output
	writer.Flush()
}
