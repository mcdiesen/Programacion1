//Funcion variatica
//suma n valores
package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}
	return total
}
func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("La suma de los 5 primeros numeros es:", total)
}

/*
run the programs: go run funcVariaticaSuma.go
Output:
	La suma de los 5 primeros numeros es: 15
*/
