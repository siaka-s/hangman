package jeu

import "strings"

type Game struct {
	Status      string   // statut du jeu à chaque tour
	lettre      []string // lettre du mot à trouver
	lettreFound []string // lettre trouvé dans le mot
	lettreUse   []string // lettre saisie
	nbrTour     int      // nombre de chance
}

func New(nbrTour int, mot string) *Game {

	letters := strings.Split(mot, "")

	found := make([]string, len(letters))

	for i := 0; i < len(letters); i++ {

		found[i] = "_"

	}

	g := &Game{
		Status:      "",
		lettre:      letters,
		lettreFound: found,
		lettreUse:   []string{},
		nbrTour:     nbrTour,
	}

	return g
}
