package main

import (
	"fmt"
	"github.com/donvito/hellomod"
)

func main() {
	fmt.Println("Modularizacion en Go!")
	//Variable: funcion que imprime los valores de las variables
	//Ubicado en el paquete unit1
	//unit1.Variable()

	//Acceso a la funcion Hola desde el paquete modularizacion
	//modularizacion.Hola()

	//geometria Acceso al paquete geometria
	/*cua1 := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cua1)

	cir1 := figuras.Circulo{Radio: 9}
	figuras.Medidas(&cir1)
	*/

	//Instancia de objeto Persona > p1
	/*p1 := models.Persona{}
	p1.Constructor("Charles", 39)
	//fmt.Println(p1)
	fmt.Println("Persona 1> Nombre: ", p1.GetNombre())
	fmt.Println("Persona 1> Edad: ", p1.GetEdad())

	//Instanciar un nuevo objeto tipo Persona >p2
	p2 := models.Persona{}
	p2.SetNombre("Charlize")
	p2.SetEdad(10)
	fmt.Println("Persona 2> Nombre:", p2.GetNombre())
	fmt.Println("Persona 2>Edad: ", p2.GetEdad())
	*/

	hellomod.SayHello()
}
