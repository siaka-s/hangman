package main

import (
	"fmt"
	"hangman/jeu"
	"os"
)

func main() {

	result := jeu.New(8, "siaka")

	choix, err := jeu.LireChoix()
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // standard dans le cas ou ca ne marche pas
	}

	fmt.Println(choix)

	fmt.Println(result)

}
