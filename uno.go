package main

import "fmt"

func sumar() uint8 {
	var number1, number2 uint8
	fmt.Println(" Digite el numero 1")
	fmt.Scanln(&number1)
	fmt.Println(" Digite el numero 2")
	fmt.Scanln(&number2)
	return number1 + number2
}

func cociente() {
	var number1, number2 uint16
	fmt.Println(" Digite el numero 1")
	fmt.Scanln(&number1)
	fmt.Println(" Digite el numero 2")
	fmt.Scanln(&number2)
	var cocientee uint16 = number1 / number2
	var residuo uint16 = number1 % number2
	fmt.Println(" El cociente es: ", cocientee, " El residuo es: ", residuo)
}

func igb() {
	var number1 float32
	fmt.Println(" Digite el valor de venta")
	fmt.Scanln(&number1)
	var igv float32 = (number1 * 18) / 100
	fmt.Printf("El impuesto es %g  y el total de compra es %g ", igv, number1+igv)
}

func main() {
	var option uint8
	fmt.Println(" Por favor digite alguna opcion del menu: \n 1- Opcion: Sumar dos numeros  \n 2- Opcion: Cociente  \n 3- Opcion: Igb ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println("El resultado  de la suma es: ", sumar())
	case 2:
		cociente()
	case 3:
		igb()
	default:
		main()
	}
}
