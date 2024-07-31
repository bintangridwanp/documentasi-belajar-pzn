package main

import "fmt"

type Handphone struct {
	NamaHandphone string
	Nomer         int64
}

func main() {
	var tipeHandphone Handphone
	fmt.Println(tipeHandphone)

	tipeHandphone.NamaHandphone = "Samsung"
	tipeHandphone.Nomer = 81575192121

	fmt.Println(tipeHandphone)
	fmt.Println(tipeHandphone.NamaHandphone)
	fmt.Println(tipeHandphone.Nomer)

	yanto := Handphone{
		NamaHandphone: "Iphone 15 Pro MAX",
		Nomer:         8157597898,
	}

	fmt.Println(yanto.NamaHandphone)
	fmt.Println(yanto.Nomer)

	agus := Handphone{"mi", 815737327887}
	fmt.Println(agus)

}
