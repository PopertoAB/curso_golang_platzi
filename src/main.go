package main

import "fmt"

func main() {
	// Declaración de contantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables
	base := 12
	var altura int = 14
	var area int

	area = base * altura

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	var baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area de un cuadrado:", areaCuadrado)
}
