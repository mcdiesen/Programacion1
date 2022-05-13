//Arreglos
package arreglos

import "fmt"

func main() {
	//Declaracion del arreglo planeta con capacidad de 8 string
	var planets [8]string

	//Asignación de valores en el arreglo - inicia en la posicion cero
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	earth := planets[2]
	//imprime Earth
	fmt.Println(earth)
}
