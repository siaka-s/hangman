package main

import (
	"fmt"
	"hangman/jeu"
)

func main() {

	result := jeu.New(8, "siaka")

	fmt.Println(result)

}
