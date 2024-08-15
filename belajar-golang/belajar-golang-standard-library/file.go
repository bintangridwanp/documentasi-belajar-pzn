package main

import (
	"bufio" // Mengimpor paket bufio untuk buffered I/O
	"io"    // Mengimpor paket io untuk operasi input/output
	"os"    // Mengimpor paket os untuk operasi sistem seperti membuka file
)

// Fungsi untuk membuat file baru dan menulis pesan ke dalamnya
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

// Fungsi untuk menambahkan pesan ke file yang sudah ada
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

// Fungsi untuk membaca isi file dan mengembalikan sebagai string
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// Membuat file baru dengan nama "sample.log" dan menulis pesan ke dalamnya
	// createNewFile("sample.log", "this is sample log")

	// Membaca isi file "sample.log" dan mencetaknya ke konsol
	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	// Menambahkan pesan ke file "sample.log"
	addToFile("sample.log", "\nthis is add message")
}
