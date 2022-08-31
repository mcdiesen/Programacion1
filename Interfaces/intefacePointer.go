//Inteface pointer
package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

//NewPerson constructor de la estructura Person
//devuelve un puntero Person
func NewPerson(name string) *Person {
	return &Person{name}
}

//Get implementa el metodo Get de la interface Storager
func (p *Person) Get() string {
	return p.name
}

//Set implementa el método Set de la interface Storager
func (p *Person) Set(name string) {
	p.name = name
}

func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("Charles")
	Exec(p, "McLean")
}
