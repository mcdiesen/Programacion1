package main

import "fmt"

//Animal interface
type animal interface {
	Run() string
	Roar() string
}

//Dog structura vacia
type Dog struct{}

//Roar implementa la interface
func (d Dog) Roar() string {
	return "woof"
}

//Run implementa el método Run
func (d Dog) Run() string {
	return "run like a Dog"
}

//Cat struct
type Cat struct{}

//Roar implementa el método Roar
func (c *Cat) Roar() string {
	return "Meow"
}
func (d *Cat) Run() string {
	return "run like a Cat"
}

//RoarAndRun de tipo animal
func RoarAndRun(a animal) {
	fmt.Printf("%s and %s\n", a.Roar(), a.Run())
}

func main() {
	myDog := Dog{}
	myCat := Cat{}

	RoarAndRun(myDog)
	//paso por referencia de la instancia myCat
	RoarAndRun(&myCat)
}

//Genera como salida del programa:
//woof and run like a Dog
//Meow and run like a Cat
