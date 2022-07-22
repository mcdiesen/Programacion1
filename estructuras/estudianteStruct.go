//Estructura Estudiante
package main

import "fmt"

type Estudiante struct {
	nombre  string
	carnet  string
	carrera string
}

func (e Estudiante) imprimir() {
	fmt.Println("Los datos del estudiante:")
	fmt.Println("Nombre: ", e.nombre, "Carnet: ", e.carnet, "Carrera:", e.carrera)
}

func (e *Estudiante) NuevoEst(nomb, carn, carre string) {
	e.nombre = nomb
	e.carnet = carn
	e.carrera = carre
}

func main() {
	//crea un objeto est1 con valores inicial
	est1 := Estudiante{nombre: "Charles", carnet: "607-XXXXX-0100", carrera: "Informática Administrativa"}
	est1.imprimir()

	//crea un objeto est2 a través del método NuevoEst y se le pasan
	//los valores
	est2 := Estudiante{}
	est2.NuevoEst("McLean", "22-002122-0100", "Ingeniería en Sistema")
	est2.imprimir()
}
