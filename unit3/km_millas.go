/*
Cree un algoritmo para convertir kilómetros en millas.
La conversión se debe llevar a
cabo mediante una función. También implemente una subrutina para
mostrar una leyenda con el resultado.
Considerar la constante FACTOR 0.6214.
*/
package main

import "fmt"

func convertirKmAM(km float64) {
	//const FACTOR = 0.6214
	milla := km * 0.62137
	imprimirResultado(milla)
}
func imprimirResultado(convertido float64) {
	fmt.Println("Conversion:", convertido)
}
func main() {
	var km float64
	fmt.Println("Ingrese el Km")
	fmt.Scanf("%f", &km)
	convertirKmAM(km)
	fmt.Println("Gracias por su tiempo!")
}
