//embeber (unir) interfaces
package main

import "fmt"

type Greeter interface {
	Greet()
}
type Byer interface {
	Bye()
}

//GreeterByer unir interfaces simples (GreeterByer)
type GreeterByer interface {
	Greeter
	Byer
}

//All Implementa los mètodos de las interfaces Greeter y Byer
func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}

}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}
func (t Text) Bye() {
	fmt.Printf("\tAdios soy %s\n", t)
}

type Persona struct {
	Name string
}

func (p Persona) Greet() {
	fmt.Printf("Hola Soy %s\n", p.Name)
}
func (p Persona) Bye() {
	fmt.Printf("\tAdios soy  %s\n", p.Name)
}

func main() {
	p := Persona{Name: "Charles"}
	var t = Text("Shania")
	All(p, t)
}
