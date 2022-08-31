//Interface animal
package main

import "fmt"

type Animal interface {
	mover() string
}
type Perro struct {
}
type Pez struct {
}
type Pajaro struct {
}

//Mover método retorna un string
func (*Perro) mover() string {
	return "Soy un perro y camino"
}
func (*Pez) mover() string {
	return "Soy un pez y nado"
}
func (*Pajaro) mover() string {
	return "Soy un pajaro y vuelo"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())

}
func main() {
	perror := Perro{}
	moverAnimal(&perror)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)

}
