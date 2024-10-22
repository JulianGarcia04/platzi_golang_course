package main

import (
	"fmt"
)

// the size parameter must be in cm
func getSquareArea (size uint) uint {
	return size * size
}

// base and height parameter must be in cm
func getTriangleArea (base, height uint) uint {
	return (base * height)/2
}

func getFullName (names ...string) []string {
	return names
}

func variablesDeclaration () {
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

func printFeatures () {
	// printLn
	fmt.Println("with jump line");
	fmt.Println("write another line");

	// printf

	language := "golang";
	platform := "platzi";

	fmt.Printf("Estoy empezando a aprender %q por %q.\n", language, platform);

	var name string = "Julian"
	var age uint8 = 20

	fmt.Printf("Hola soy %s y tengo %d años.\n", name, age);

	fmt.Printf("Esta variable es de tipo %T\n", age);

	// Sprintf

	var time uint16 = 60

	var message string = fmt.Sprintf("Esta variable se almacenará en tiempo prolongado en minutos. Tiempo: %d\n", time);

	fmt.Print(message)
}

func functions () {
	squareArea := getSquareArea(10);
	fmt.Printf("El area del cuadrado es %d centimetros.\n", squareArea);

	triangleArea := getTriangleArea(20, 30);
	fmt.Printf("El area del triangulo es %d centimetros.\n", triangleArea);

	names := getFullName("Julian", "Garcia")

	fmt.Println(names)

}

func loops () {
	// for loop
	fmt.Println("for loop")

	for i := 100; i >= 0; i-- {
		fmt.Printf("número %d \n", i)
	}

	// for while loop
	fmt.Println("For while loop")

	counter := 100

	for counter >= 0 {
		fmt.Printf("el counter está en %d \n", counter);
		counter--
	}

	// for forever loop
	fmt.Println("For forever loop")

	fmt.Println("Lo comenté para ahorar recursos")

	// counterForever := 0

	// for {
	// 	fmt.Printf("el counter del for forever está en %d \n", counterForever)
	// 	counterForever++
	// }
}

func ifConditional (testValue int) {
	nro := testValue
	if nro >= 0 {
		fmt.Println("Este numero está entre cero y más allá")
		return
	}

	if nro < 0 {
		fmt.Println("Este número es negativo")
		return
	}

	fmt.Println("El número es nil o no se que es")
}

func switchConditional (nroTest int) {

	switch nro := nroTest; nro  {
	case 10:
		fmt.Println("Es un número 10")
	case 20:
		fmt.Println("Es un número 20")
	case -1:
		fmt.Println("Es un número -1")
	default:
		fmt.Println("Puede ser un número")
	}

	switch nro := nroTest; {
	case nro >= 0:
		fmt.Println("Este numero está entre cero y más allá")
	case nro < 0: 
		fmt.Println("Este número es negativo")
	default:
		fmt.Println("Es un número ???")
	}
}

func main () {
	fmt.Println("-------------hello world------------------")
	fmt.Println("Hello world!!")

	fmt.Println("-------------variables declaration-------------------")
	variablesDeclaration()

	fmt.Println("-------------print features-----------------------")
	printFeatures()

	fmt.Println("-------------functions-----------------")
	functions()

	fmt.Println("-------------loops---------------")
	loops()

	fmt.Println("-------------if conditional---------------")
	ifConditional(0)
	ifConditional(-2)

	fmt.Println("-------------switch conditional-------------")
	switchConditional(10)
	switchConditional(20)
	switchConditional(-1)

	// my first server in go :)
	// this it has the course but I wanted create a server

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)

	// err := http.ListenAndServe(":3000", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("Server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error stating server: %s\n", err);
	// 	os.Exit(1)
	// }
}