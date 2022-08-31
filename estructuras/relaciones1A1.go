//relacion entre estructuras  1 a 1
package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {
	alex := User{
		"Alex",
		"alex@gmail.com",
		true,
	}

	roel := User{
		"Roel",
		"roel@gmail.com",
		true,
	}

	alexStudent := Student{
		user:   alex, //manejo de la relacion para acceder a la relacion 1 a 1 struct
		codigo: "607-241192-0000L",
	}
	fmt.Println(roel)
	fmt.Println(alex)
	fmt.Println("Nombre estudiante: ", alexStudent.user.nombre, "Codigo estudiante ", alexStudent.codigo)

}
