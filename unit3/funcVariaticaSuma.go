//Funcion variatica
//suma n valores
package main

import "fmt"

//Sum funcion variatica que recibe los numeros a sumar
func sum(nums ...int) (int, int) {
	total := 0
	var I int
	I = len(nums)
	for _, num := range nums {
		total += num
	}
	return I, total
}
func main() {
	nums, total := sum(1, 2, 3, 4, 5, 7)
	fmt.Println("La suma de los ", nums, "numeros ingresados es:", total)
}

/*
	La suma de los 6 numeros ingresados es: 22
*/
