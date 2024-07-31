package main

import "fmt"

func main(){
	 var nilai32  int32 = 2000000
		nilai64 := int64(nilai32)
		nilai8 := int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama ="Bintang"
	 b :=  nama[3]
	 bString := string(b)

	fmt.Println(nama)
	fmt.Println(bString)

}