package main

import "fmt"

func main() {

	var num1 int
	var num2 int

	fmt.Println("Ingrese el primer numero")
	fmt.Scan(&num1)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("El numero mayor es: ", num1)
	} else if num1 == num2 {
		fmt.Println("Ambos numeros son iguales")
	} else {
		fmt.Println("El numero mayor es: ", num2)
	}

}
