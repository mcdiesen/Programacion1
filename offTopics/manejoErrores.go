//Manejo de errores
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Musketeers = []string{
	"Athos", "POrthos", "Aramis", "D'Artagnan",
}

func GetMusketeer(id int) (string, error) {
	if id < 0 || id > len(Musketeers) {
		return "", error.New(
			fmt.Sprintf("Invalid id [%d]", id))
	}
	return Musketeers[id], nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int() % 6
	mosq, err := GetMusketeer(id)

	if err == nil {
		fmt.Println("[%d] %s", id, mosq)
	} else {
		fmt.Println(err)
	}
}
