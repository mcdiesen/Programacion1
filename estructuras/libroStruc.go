package main

import "fmt"

type Libro struct {
	titulo  string
	autor   string
	precio  float32
	edicion int
}

type Estudiante1 struct {
	nombre         string
	edad           int32
	carnet         string
	cod_carrera    string
	nombre_carrera string
	Libro          []string
}

//Asignar metodo que crear libros
func (l Libro) Asignar() {

}

//Presar metodo que permite prestar libro
func (l *Libro) Prestar() {

}

//CargarEstudiante metodo que permite crear un estudiante
func (e Estudiante) CargarEstudiante() {

}

//MostrarEstudiante metodo que muestra el estudiante
func (e Estudiante) MostrarEstudiante() {

}

func main() {
	libro1 := Libro{}
	fmt.Println(libro1)
}
