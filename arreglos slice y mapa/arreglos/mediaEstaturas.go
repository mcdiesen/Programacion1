//Calcula la media de estaturas de una clase
//Deduce cuantos son ms altos que la media y cuantos son mas bajos que dicha media.
package main

import "fmt"

func main() {
	var N int
	fmt.Println("Cuantos estudiantes tiene la clase? ")
	fmt.Scan(&N)

	NumEst := make([]float32, 0.0)
	//var Suma float32 = 0.0

	fmt.Println("Ingrese las estaturas de los estudiantes: ")

	for x := 0; x < N; x++ {
		var estatura float32
		fmt.Println("Estatura: ", x+1)
		fmt.Scanf("%f", &estatura)
		NumEst = append(NumEst, estatura)
	}
	fmt.Println("Estaturas ingresadas: ", NumEst)

}
