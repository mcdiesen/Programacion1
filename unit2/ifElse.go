package main

import "fmt"

//Diseñar un programa que lea tres números y descubra si uno de ellos es la suma de los otros dos.
func main() {
	var num1, num2, num3 int

	fmt.Println("Ingrese los valores para numero 1, numero 2 y numero 3 respectivamente:")
	fmt.Scan(&num1, &num2, &num3)

	if (num1 + num2) == num3 {
		fmt.Println(num1, "+", num2, "es == a ", num3)
	} else if num2+num3 == num1 {
		fmt.Println(num2, "+", num3, "es == a ", num1)
	} else {
		fmt.Println("No son iguales")
	}
}
