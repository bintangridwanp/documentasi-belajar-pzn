package main

import "fmt"

func main() {

	//tipe Data array
	//jika tadak mau menentukan nilai array di awal bisa di ganti dengan [...]
	var names [4]string

	names[0] = "bintang"
	names[1] = "ridwan"
	names[2] = "pribadi"
	names[3] = "keren"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	//membuat array langsung

	var values = [...]int{
		10,
		20,
		30,
	}

	fmt.Println(values)

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//len untuk mengecek panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))

}
