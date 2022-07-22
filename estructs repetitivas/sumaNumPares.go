//Realizar un programa para obtener la suma de los numeros pares
//hasta 1000 inclusive
package main

import "fmt"

func main() {
	numero := 2
	suma := 0

	for i := 0; i <= 1000; i++ {
		suma = suma + numero
		numero += 2
	}
	fmt.Println("La suma = ", suma)
}
