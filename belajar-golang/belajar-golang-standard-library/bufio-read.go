package main

import (
	"bufio"   // Mengimpor paket bufio untuk buffered I/O
	"fmt"     // Mengimpor paket fmt untuk mencetak output ke konsol
	"io"      // Mengimpor paket io untuk operasi input/output
	"strings" // Mengimpor paket strings untuk memanipulasi string
)

func main() {
	// Membuat pembaca string dari input string
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	// Membaca setiap baris dari input dalam loop
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			// Jika mencapai akhir file, keluar dari loop
			break
		}

		// Mencetak setiap baris yang dibaca
		fmt.Println(string(line))
	}
}
