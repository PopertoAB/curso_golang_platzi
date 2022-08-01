package main

import (
	"fmt"
	"math"
)

func main() {

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)
	
	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)
	
	// Módulo
	result = y % x
	fmt.Println("Módulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)
	
	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// -Rectángulo, trapecio y círculo
	const base = 10
	const altura = 20
	area := base * altura
	fmt.Println("El area de un rectángulo es de:", area)

	const baseMayor = 10
	const baseMenor = 5
	const alturaTrapecio = 3
	areaTrapecio := ( ( baseMayor + baseMenor ) / 2.0 ) * alturaTrapecio
	fmt.Println("El area de un trapecio es de:", areaTrapecio)

	const radio = 10
	areaCirculo := math.Pi * radio * radio
	fmt.Println("El area un un circulo es de:", areaCirculo)
}
