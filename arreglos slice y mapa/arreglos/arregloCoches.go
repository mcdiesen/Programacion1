//arreglo Coches
package main

import "fmt"

func main() {
	//arreglo Coches de nueve elementos
	var Coches = [9]string{
		"Alfa Romeo",
		"Fiat",
		"Ford",
		"Lancia",
		"Renault",
		"Seat",
	}
	fmt.Println(Coches)
	n := cap(Coches)
	//fmt.Println(n)
	for i := 0; i < n; i++ {
		if Coches[i] == "Lancia" {
			posic1 := Coches[i]

		}
	}
	fmt.Println(Coches)

}
