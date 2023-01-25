package main

import "fmt"

func main() {
	fmt.Println("Resolução Parte 1:")

	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Print(x, "-")
		}
	}

	fmt.Println("\n\n")
	fmt.Println("Resolução Parte 2:")

	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("Pin-Pan")
		} else if x%3 == 0 && x%5 != 0 {
			fmt.Println("Pin")
		} else if x%3 != 0 && x%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(x)
		}
	}

}
