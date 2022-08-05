package main

import (
	"fmt"
	"time"
)

type Fecha struct {
	dia  int
	mes  int
	anio int
}
type Persona struct {
	nombre          string
	apellidos       string
	fechaNacimiento Fecha
}

func (f Persona) Edad(d int, a int) {
	//	fechaNac := 2022 - a
	//fmt.Println("Año de nacimiento", a, "Tienes ", fechaNac, "años")
	aa := time.Date(a, 11, d, 0, 0, 0, 0, time.UTC)
	anioNac, _, _ := aa.Date()

	fechaNac := 2022 - anioNac
	fmt.Println(f.nombre, "tienes", fechaNac, "años de edad")
}
func (p Persona) Persona(unNombre string, unApp string, unaFecha Fecha) {
	p.nombre = unNombre
	p.apellidos = unApp
	p.fechaNacimiento = unaFecha
}

func main() {
	unaPersona := Persona{"Charles", "McLean", Fecha{24, 11, 1982}}

	unaPersona.Edad(24, 1982)
	fmt.Println(unaPersona)
}
