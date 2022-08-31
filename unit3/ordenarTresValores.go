/*
Realizar un algoritmo que permita ordenar tres números
mediante un procedimiento de intercambio en dos variables
(paso de parámetros por referencia).
*/
package main

import "fmt"

func intercambio(a, b *float64) {
	var auxi float64
	auxi = *a
	*a = *b
	*b = auxi
}
func main() {
	var x, y, z float64
	fmt.Println("Deme tres numeros reales")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	if x > y {
		intercambio(&x, &y)
	}
	if y > z {
		intercambio(&y, &z)
	}
	if x > y {
		intercambio(&x, &y)
	}
	fmt.Println(x, y, z)
}
