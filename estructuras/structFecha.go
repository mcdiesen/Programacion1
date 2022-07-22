package main

import "fmt"

type Fecha struct {
	dia, mes, anio int
}

func (f Fecha) MostrarFecha() {
	fmt.Println("La fecha es: ", f.dia, "/", f.mes, "/", f.anio)
}
func (f *Fecha) AsignarFecha(d, m, a int) {
	f.dia = d
	f.mes = m
	f.anio = a
}

func main() {
	f1 := Fecha{}
	f1.AsignarFecha(5, 7, 2022)
	f1.MostrarFecha()
}
