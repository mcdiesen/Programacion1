package main

import "fmt"

type Fechas struct {
	dia  int
	mes  int
	anio int
}

func (f Fechas) Compara() {
	fmt.Println("la fecha es:", f.dia, f.mes, f.anio)
}
func main() {
	f1 := Fechas{4, 8, 2022}
	f1.Compara()
}
