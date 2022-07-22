package main

import "fmt"

type Cadena struct {
	cad string
}

//Mostrar metodo de la estructura Cadena que muestra por pantalla el valor de la
//propiedad cad
func (c Cadena) Mostrar() {
	fmt.Println(c.cad)
	//copy
}

//Longitud de la cadena
func (c Cadena) longitud() {
	long := len(c.cad)
	fmt.Println("La longitud de la cadena es: ", long)
}
func main() {
	//miCadena instancia de la estructura Cadena
	miCadena := Cadena{cad: "Hola Mundo!"}

	fmt.Println(miCadena)
	miCadena.longitud()

	miCadena.cad = "Hola Cadena..."

	//miCadena.Mostrar() invoca al método Mostrar que imprime el valor de la
	//propiedad cad
	miCadena.Mostrar()
	miCadena.longitud()
}
