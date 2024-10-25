package variables_declaration

import "fmt"

func VariablesDeclaration () {
	// Declaración de variables
	const pi float64 = 3.14

	const pi2 = 3.1415

	fmt.Println("pi ", pi)
	fmt.Println("pi2", pi2)

	// Declaración de variables enteras
	// Primera forma
	base := 12
	// Segunda forma
	var altura int = 14
	// Tercera forma
	var area int

	fmt.Println("base ", base)
	fmt.Println("altura ", altura)
	fmt.Println("area ", area)

	// Zero values
	var number int
	var float float64
	var string string
	var boolean bool

	fmt.Println("zero value of number variable ", number)
	fmt.Println("zero value of float variable ", float)
	fmt.Println("zero value of string variable ", string)
	fmt.Println("zero value of boolean variable ", boolean)

	// Example
	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("area del cuadrado ", areaCuadrado)
}