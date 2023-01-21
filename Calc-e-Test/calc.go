package main

import "fmt"

func main() {

	fmt.Println(somar(1, 1, 0.5))       // todos os números são somados
	fmt.Println(subtrair(5, 1, 1, 0.5)) // Do primeiro número serão subtraídos todos os demais
	fmt.Println(multiplicar(2, 1.25))   // todos os números são multiplicados
	fmt.Println(dividir(100, 4, 5, 2))  // O primeiro número será dividido pelo próximo, e o resultado será dividido pelo próximo, até o último
}

func somar(numeros ...float64) float64 {
	total := 0.0
	for _, v := range numeros {
		total += v
	}
	return total
}

func subtrair(numero float64, numeros ...float64) float64 {
	total := numero
	for _, v := range numeros {
		total -= v
	}
	return total
}

func multiplicar(numeros ...float64) float64 {
	total := 1.0
	for _, v := range numeros {
		total = total * v
	}
	return total
}

func dividir(numero float64, numeros ...float64) float64 {
	total := numero
	for _, v := range numeros {
		total = total / v
	}
	return total
}
