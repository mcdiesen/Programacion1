//Lectura de 20 elementos de un arreglo
package main

import "fmt"

func main() {
	var Arreglo1 [20]int
	for itm := range Arreglo1 {
		fmt.Scan(&Arreglo1[itm])
	}
	fmt.Println("Los elementos del arreglo son: ")
	for itm := range Arreglo1 {
		fmt.Printf("%d \t", Arreglo1[itm])
	}
}
