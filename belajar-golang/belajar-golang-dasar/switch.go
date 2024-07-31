package main

import "fmt"

func main() {
	buah := "jerukkkkkkkkkkkk"

	switch buah {
	case "jeruk" :
		fmt.Println("saya jeruk")
	case "mangga":
		fmt.Println("saya mangga")
	default:
		fmt.Println("saya bukan buah-buahan")
	}

	switch lenght := len(buah); lenght > 3{
	case true:
		fmt.Println("ini buah-buahan")
	case false:
		fmt.Println("ini bukan buah-buahan")
	}

	lenght := len(buah)
	switch{
	case lenght > 10:
		fmt.Println("ini buah-buahan")
	case lenght > 5:
		fmt.Println("ini buah")
	default:
		fmt.Println("ini bukan buah-buahan")
	}
}

