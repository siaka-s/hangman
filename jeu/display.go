package jeu

import (
	"fmt"
)

func Dessin() {

	fmt.Println(`
     ##    ##   #  ##   ##           ##   ##   ####    ###  ##   ##  ##  ### ###   ####    ###  ##
     ##    ## #    ##   ##           ##   ##  ##  ##   #### ##  ##       #######  ##  ##   #### ##
     ##    ####    ##   ##           #######  ##  ##   ## ####  ##       #######  ##  ##   ## ####
 ##  ##    ## #    ##   ##           ##   ##  ######   ##  ###  ##  ###  ## # ##  ######   ##  ###
 ##  ##    ##   #  ##   ##           ##   ##  ##  ##   ##   ##   ##  ##  ##   ##  ##  ##   ##   ##
  ####    #######   #####            ##   ##  ##  ##   ##   ##    #####  ##   ##  ##  ##   ##   ##
`)
}

func DessinEtat(g *Game, choix string) {
	DessinTour(g.nbrTour)
	DessinStatut(g, choix)
}

func DessinTour(tour int) {

	var etats = [8]string{
		`
  -----
  |   |
      |
      |
      |
      |
=========
`, // 0 essais

		`
  -----
  |   |
  O   |
      |
      |
      |
=========
`, // 1 essai

		`
  -----
  |   |
  O   |
  |   |
      |
      |
=========
`, // 2 essais

		`
  -----
  |   |
  O   |
 /|   |
      |
      |
=========
`, // 3 essais

		`
  -----
  |   |
  O   |
 /|\  |
      |
      |
=========
`, // 4 essais

		`
  -----
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`, // 5 essais

		`
  -----
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`, // 6 essais

		`
  -----
  |   |
  O   |
 /|\  |
 / \  |
YOU WIN!
=========
`, // 7 essais - Partie perdu
	}

	switch tour {
	case 0:
		fmt.Println(etats[tour-1])
	case 1:
		fmt.Println(etats[tour-2])
	case 2:
		fmt.Println(etats[tour-3])
	case 3:
		fmt.Println(etats[tour-4])
	case 4:
		fmt.Println(etats[tour-5])
	case 5:
		fmt.Println(etats[tour-6])
	case 6:
		fmt.Println(etats[tour-7])
	case 7:
		fmt.Println(etats[tour-8])
	}

}

func DessinStatut(g *Game, choix string) {

	fmt.Printf("Trouvé :")

	construireLettre(g.lettreFound)

	fmt.Println("Utilisé :")

	construireLettre(g.lettreUse)

	switch g.Status {
	case "Bon choix":
		fmt.Println("voius avez choisi une bonne lettre")
	case "deja saisi":
		fmt.Printf("vous avez deja saisi cette lettre : %v", choix)
	case "Mauvais choix":
		fmt.Printf("vous avez choisi une mauvaise lettre  : %v", choix)
	case "perdu":
		fmt.Printf("Vous avez perdu")
		construireLettre(g.lettre)
	case "gagné":
		fmt.Printf("Vous avez gagné")
		construireLettre(g.lettre)

	}

}

func construireLettre(l []string) { // contruire les lettres trouvé et utiulisé a chaque essaie

	for _, c := range l {

		fmt.Printf("%v_", c)
	}
	fmt.Println()

}
