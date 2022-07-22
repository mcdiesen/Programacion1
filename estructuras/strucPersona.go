package main

import "fmt"

type empleado struct {
	segSocial string
}

func (e empleado) setSegSocial(segSoc string) string {
	e.segSocial = segSoc

	return "se pudo hacer el registro"
}
func (e empleado) getSegSocial() {
	fmt.Println(e.segSocial)
}

type persona struct {
	nombre string
	emp    empleado
}

func (p persona) walk() {
	fmt.Println("Persona", p.nombre, "Caminando")
}

func main() {
	p1 := persona{"Charles", empleado{"123", "325"}}
	p1.walk()
	p1.emp.setSegSocial("999")
	p1.emp.getSegSocial()

}
