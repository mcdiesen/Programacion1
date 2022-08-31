//Interface Greeter que es implementado por la struct Person y rl tipo Text
package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

//Text tipo string
type Text string

//Greet() implementado por el tipo Text
func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Adios soy %s\n", t)
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s\n", p.Name)
}
func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
	}
}

//ByeAll función variatica que implementa la Byer
func ByeAll(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
	}
}
func main() {
	p := Person{Name: "Charles"}
	t := Text("Shania")
	GreetAll(p, t)
	ByeAll(p, t)
}
