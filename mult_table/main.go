package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("a\\b\t")

	for _, number := range numbers {
		fmt.Printf("%d\t", number)
	}
	fmt.Println()
	for _, i := range numbers {
		fmt.Printf("%d\t", i)
		for _, j := range numbers {
			fmt.Printf("%d\t", i*j)
		}

		fmt.Println()
	}
}
