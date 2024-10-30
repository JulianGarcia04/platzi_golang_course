package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/JulianGarcia04/platzi_golang_course/src/concurrency"
	conditionalsStructure "github.com/JulianGarcia04/platzi_golang_course/src/conditionals_structures"
	fmtPrinters "github.com/JulianGarcia04/platzi_golang_course/src/fmt_print_features"
	"github.com/JulianGarcia04/platzi_golang_course/src/functions"
	variablesDeclaration "github.com/JulianGarcia04/platzi_golang_course/src/variables_declaration"
)

type User struct {
	username string;
	firstname string;
	lastname string;
	age int;
	createAt string;
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

func arrays () {
	numberArray := [4]int{10, 20, 30, 40}

	fmt.Printf("array:%v; length: %d; capacity: %d; \n", numberArray, len(numberArray), cap(numberArray))

	fmt.Println(numberArray[0])
	fmt.Println(numberArray[:2])
	fmt.Println(numberArray[1:2])

	numberArray[0] = 15

	fmt.Println(numberArray[0])
}

func slicesFunc () {
	numberSlices := []int{10, 20, 30, 40, 50, 60}

	fmt.Printf("slice: %v; length: %d; capacity: %d; \n", numberSlices, len(numberSlices), cap(numberSlices));

	fmt.Println(numberSlices[0])
	fmt.Println(numberSlices[:2])
	fmt.Println(numberSlices[1:2])

	// append slices

	numberSlices = append(numberSlices, 70, 80)

	fmt.Printf("slice: %v; length: %d; capacity: %d; \n", numberSlices, len(numberSlices), cap(numberSlices));

	numberSlices = append(numberSlices, []int{90, 100, 110, 120}...)

	fmt.Printf("slice: %v; length: %d; capacity: %d; \n", numberSlices, len(numberSlices), cap(numberSlices));
}

func rangeMethod () {
	names := []string{"julian", "rosario", "maria", "emma", "juan", "orlando", "fernando"}

	fmt.Println("--------only values----------")
	for _, valor := range names {
		fmt.Println(valor)
	}
	fmt.Println("-------only index--------")
	for i := range names {
		fmt.Println(i)
	}
	fmt.Println("--------index and values----------")
	for i, name := range names {
		fmt.Println(i, name)
	}
}

func mapsFunc () {
	m := map[string]int{
		"Jose": 15,
		"Julian": 20,
		"Hernando": 30,
	}

	fmt.Println(m)

	// recorrer
	fmt.Println("go over map ")
	for k, v := range m {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	// seleccionar valor
	fmt.Println("Select a value")
	name := "Jose"

	age := m[name]

	fmt.Printf("%s tiene %d años\n", name, age)
}

func implementStructs () {
	var addUser = User{username: "julian04", firstname: "julian", lastname: "garcia", age: 20, createAt: "13/02/2004"};

	log.Println(addUser)

	users := []User{ {username: "Fer", firstname: "Fernanda", lastname: "Sandoval", age: 22, createAt: "20/12/2003"} };

	users = append(users, addUser)

	log.Println(users)

	checkIfExistUser := func (E User) bool {
		return E.username == "julian04"
	}

	log.Println(slices.ContainsFunc(users, checkIfExistUser))

	log.Println(users[0])
}

func main () {
	fmt.Println("-------------hello world------------------")
	fmt.Println("Hello world!!")

	fmt.Println("-------------variables declaration-------------------")
	variablesDeclaration.VariablesDeclaration()

	fmt.Println("-------------print features-----------------------")
	fmtPrinters.PrintFeatures()

	fmt.Println("-------------functions-----------------")
	functions.Functions()

	fmt.Println("-------------loops---------------")
	loops()

	fmt.Println("-------------if conditional---------------")
	conditionalsStructure.IfConditional(0)
	conditionalsStructure.IfConditional(-2)

	fmt.Println("----------------switch conditional-------------")
	conditionalsStructure.SwitchConditional(10)
	conditionalsStructure.SwitchConditional(20)
	conditionalsStructure.SwitchConditional(-1)

	fmt.Println("------------------arrays--------------------")
	arrays()

	fmt.Println("------------------slices-----------------------")
	slicesFunc()

	fmt.Println("------------------range method-----------------------")
	rangeMethod()

	fmt.Println("------------------is palindrome function-----------------------")
	fmt.Printf("'rice' is palidrome? %v \n", functions.IsPalindrome("rice"))
	fmt.Printf("'aibofobia' is palidrome? %v \n", functions.IsPalindrome("aibofobia"))
	fmt.Printf("'Aibofobia' is palidrome? %v \n", functions.IsPalindrome("Aibofobia"))

	fmt.Println("-----------------maps-------------------")
	mapsFunc()

	fmt.Println("-----------------strucs------------------")
	implementStructs()

	fmt.Println("-----------------concurrency------------------")
	concurrency.Concurrency()


	// my first server in go :)
	// this it doesn't has the course but I wanted create a server

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