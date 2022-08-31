//Aplicar una relacion 1 a muchos
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

//Relacion uno a muchos
type Curso struct {
	titulo string
	videos []Video
}
type Video struct {
	titulo string
	curso  Curso
}

func main() {
	//relacion 1 a 1
	/*	alex := User{nombre: "Charles",
		email:  "mcdies@gmail.com",
		activo: true,
	}*/
	//relacion 1 a muchos
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalación"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}
	video1.curso = curso
	video2.curso = curso
	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
