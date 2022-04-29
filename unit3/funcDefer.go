package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don’t just check errors, handle them gracefully.")
	return err
}
func main() {
	//El archivo que creará se llama hola.txt
	pv := proverbs("hola.txt")
	fmt.Println("Archivo creado con éxito", pv)
}
