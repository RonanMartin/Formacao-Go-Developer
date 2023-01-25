package main

import "fmt"

const ebulicaoK = 373.15

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273.15 // Conversão de Kelvin para Celsius, apenas subtrai 273.15

	fmt.Printf("A temperatura de ebulição da água em °K = %g e a temperatura de ebulição da água em °C = %g.", tempK, tempC)
}
