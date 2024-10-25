package fmt_print_features

import "fmt"

func PrintFeatures () {
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