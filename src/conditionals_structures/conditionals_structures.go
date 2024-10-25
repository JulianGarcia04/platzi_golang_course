package conditionals_structures

import "fmt"

func IfConditional (testValue int) {
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

func SwitchConditional (nroTest int) {

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