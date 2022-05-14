package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	fmt.Printf("Estas avanzando %s en la programación en Go\n", name)
}
