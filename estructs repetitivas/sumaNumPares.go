//Realizar un programa para obtener la suma de los numeros pares
//hasta 10 inclusive
package main

import "fmt"

func main() {
	numero := 2
	suma := 0

	for i := 0; i <= 5; i++ {
		suma = suma + numero
		numero += 2
	}
	fmt.Println("La suma = ", suma)
}
