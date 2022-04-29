package main

import "fmt"

func main() {
	var ptr1 *int
	var numeros int = 10
	ptr1 = &numeros
	//modifica el valor de la variable numeros
	*ptr1 = 20
	//imprime la dirección de memoria y el valor de la variable numeros
	fmt.Println("La dirección de memoria de numeros:", ptr1)
	fmt.Println("El valor almacenado en numeros es:", numeros)
}
