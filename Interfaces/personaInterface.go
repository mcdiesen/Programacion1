package main

import "fmt"

//Greeter interface
type Greeter interface {
	SayHello() string
}

//Person struct
type Person struct {
	name string
}

//SayHello implementado por la estructura Person
func (p *Person) SayHello() string {
	return "Hello, This is me " + p.name
}

func main() {
	//g de tipo interface Greeter
	var g Greeter
	p := Person{"Charles"}

	g = &p

	fmt.Println(g.SayHello())
}
