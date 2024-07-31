package main

import "fmt"

func main(){

	//Operasi Perbandingan String
	 nama1 := "bintang"
	 nama2 := "bintang"

	var result bool = nama1 == nama2

	fmt.Println(result)

	//Operasi Perbandingan int

	var value1 = 20
	var value2 = 10

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}