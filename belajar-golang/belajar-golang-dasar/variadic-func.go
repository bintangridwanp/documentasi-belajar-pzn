package main

import "fmt"

func penjumlahan(numbers ...int) int {
	total := 0

	for _, number := range numbers{
		total += number
	}

	return total
}

func main() {
	fmt.Println(penjumlahan(10, 10, 10, 10))
	fmt.Println( penjumlahan(10, 10, 10, 10, 10, 10, 10))
	fmt.Println( penjumlahan(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10}
	fmt.Println(penjumlahan(numbers...))

}