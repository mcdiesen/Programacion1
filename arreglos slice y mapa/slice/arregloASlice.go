package main

import "fmt"

func main() {
	//declarar un arreglo llamado dwarfArray
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	//concierte el arreglo a un slice
	dwarfSlice := dwarfArray[:]

	//Uso del verbo %T para comparar el arreglo y el slice
	fmt.Printf("El arreglo dwarfArray es de tipo: %T\n", dwarfArray)
	fmt.Printf("El slice dwarfSlice es de tipo: %T\n", dwarfSlice)

}
