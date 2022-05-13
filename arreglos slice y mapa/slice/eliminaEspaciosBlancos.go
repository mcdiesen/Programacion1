package main

import (
	"fmt"
	"strings"
)

// hyperspace elimina los espacios en blanco del arreglo worlds
func hyperspace(worlds []string) {
	/*This argument is a slice, not an array.*/
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{" Venus ", "         Earth        ", "       Mars"}
	hyperspace(planets)

	fmt.Println(strings.Join(planets, ""))
}
