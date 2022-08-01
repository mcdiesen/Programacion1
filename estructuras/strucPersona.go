package main

import "fmt"

type Empleado struct {
	SegSocial string
}

func (e Empleado) setSegSocial(segSoc string) string {
	e.SegSocial = segSoc

	return "se pudo hacer el registro"
}
func (e Empleado) getSegSocial() {
	fmt.Println(e.SegSocial)
}

type persona struct {
	Nombre string
	Emp    Empleado
}

func (p persona) walk() {
	fmt.Println("Persona", p.Nombre, "Caminando")
}

func main() {
	p1 := persona{"Charles", Empleado{"325"}}
	p1.walk()
	p1.Emp.setSegSocial("999")
	p1.Emp.getSegSocial()
}
