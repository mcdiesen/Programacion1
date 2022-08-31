package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

//metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\n Edad: %d\n", p.nombre, p.edad)
}
func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//Herencia la estructura Empleado hereda de Persona sus atributos y métodos
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	//Objeto persona
	p1 := Persona{"Alex", 26}
	p1.imprimir()
	p1.editNombre("Charles")
	p1.imprimir()

	//Objeto Empleado
	emp1 := Empleado{sueldo: 500}
	emp1.nombre = "Pedro"
	emp1.edad = 30
	emp1.imprimir()
	fmt.Println(emp1)
}
