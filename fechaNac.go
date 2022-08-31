package main

import (
	"fmt"
	"time"
)

func main() {
	// Fecha actual con hora actual
	currentTime := time.Now()

	// Ubicación, zona horaria
	loc := currentTime.Location()

	//La fecha pasada que tienes
	pasttimestr := "1982-11-24 00:00"

	// El formato que le decimos que use para leer las fechas
	layout := "2022-08-02 15:04"

	//Convertir tu fecha pasada a date
	pasttime, err := time.ParseInLocation(layout, pasttimestr, loc)
	if err != nil {
		fmt.Println(err)
	}

	// Calculas la diferencia
	diff := pasttime.Sub(currentTime)

	// Informas el resultado en minutos
	fmt.Printf("La diferencia es %v, o %v en minutos\n", diff, diff.Minutes())
}
