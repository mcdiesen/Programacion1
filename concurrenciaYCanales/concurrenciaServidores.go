package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(server string, canal chan string) {
	_, err := http.Get(server)
	if err == nil {
		canal <- server + " No esta disponible"
	} else {
		canal <- server + " Esta funcionando"
	}
}
func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"oregoom.com/",
		"udemy.com/",
		"youtube.com/",
		"facebook.com/",
		"google.com/",
	}
	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución ", tiempoPaso)
}
