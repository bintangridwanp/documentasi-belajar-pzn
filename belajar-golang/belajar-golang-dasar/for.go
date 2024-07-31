package main

import "fmt"

func main() {

	//counter := 1

	//for counter <= 10{
	//	fmt.Println("perulangan ke", counter)
	//	counter++
	//}
	//fmt.Println("selesai")

	for counter := 1; counter <= 10; counter++{
		fmt.Println("perulangan ke", counter)
	}
	fmt.Println("selesai")

	 nama := []string{"bintang", "ridwan", "pribadi"}
	for i := 0; i < len(nama); i++{
		fmt.Println(nama[i])
	}
	
	for index, name := range nama{
		fmt.Println("index", index, "=", name)
	}
}
