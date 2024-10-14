package main

import (
	"fmt"
	"hangman/jeu"
	"os"
)

func main() {

	jeu.Dessin()
	result := jeu.New(8, "siaka")

	choixl := ""

	for {
		jeu.DessinStatut(result, choixl)

		switch result.Status {

		case "perdu", "gagné":

			os.Exit(0)
		}
		choix, err := jeu.LireChoix()
		if err != nil {
			fmt.Println(err)
			os.Exit(1) // standard dans le cas ou ca ne marche pas
		}
		choixl = choix
	}

}
