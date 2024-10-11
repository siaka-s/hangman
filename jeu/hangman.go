package jeu

import "strings"

type Game struct {
	status      string   // statut du jeu à cahque tour
	lettre      []string // lettre du mot àn trouvé
	lettreFound []string // lettre trouvé dans le mot
	lettreUse   []string // lettre saisi
	nbrTour     int      // nombre de chance
}

func New(nbrTour int, mot string) *Game {

	letters := strings.Split(mot, "")

	found := make([]string, len(letters))

	for i := 0; i < len(letters); i++ {

		found[i] = "_"

	}

	g := &Game{
		status:      "",
		lettre:      letters,
		lettreFound: found,
		lettreUse:   []string{},
		nbrTour:     nbrTour,
	}

	return g
}
