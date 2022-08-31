//Implementar el método string de la intreface Stringer del paquete fmt
package main

import "fmt"

type Persona struct {
	Name string
}

//String implementa el método String del paquete fmt
func (p Persona) String() string {
	return "Hola soy una persona y me llamo: " + p.Name
}

func main() {
	p := Persona{"Charles"}
	fmt.Println(p)
}
