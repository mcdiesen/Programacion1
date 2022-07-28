//Calcula la media de estaturas de una clase
//Deduce cuantos son mas altos que la media y cuantos son mas bajos que dicha media.
package main

import "fmt"

func main() {
	var N, Bajos, Altos int

	fmt.Println("Cuantos estudiantes tiene la clase? ")
	fmt.Scanln(&N)

	NumEst := make([]float32, 0.0)
	var Suma, Media float32

	fmt.Println("Ingrese las estaturas de los estudiantes: ")

	for x := 0; x < N; x++ {
		var estatura float32
		fmt.Println("Estatura: ", x+1)
		fmt.Scanf("%f", &estatura)
		NumEst = append(NumEst, estatura)

		Suma = Suma + NumEst[x]
	}

	Media = Suma / float32(N)

	//Evalua cuantos estudiantes estan por debajo y por arriba la media de estaturas
	for i := 0; i < N; i++ {
		if NumEst[i] < Media {
			Bajos = Bajos + 1
		}
		if NumEst[i] > Media {
			Altos = Altos + 1
		}
	}
	fmt.Println("Estaturas ingresadas: ", NumEst)
	fmt.Println("La suma de las estaturas es: ", Suma)
	fmt.Println("La estatura media es: ", Media)
	fmt.Println("Las cantidad de estaturas por debajo de la media son: ", Bajos)
	fmt.Println("La cantidad de estaturas por arriba de la media son", Altos)

}
