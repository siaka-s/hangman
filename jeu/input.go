package jeu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lecteur = bufio.NewReader(os.Stdin) // effectuer un lecture de saisi standard

func LireChoix() (choix string, err error) {

	valid := false

	for !valid {

		fmt.Print("saisisez votre lettre svp : \n---> ")

		choix, err = lecteur.ReadString('\n') // permet de lire la saisi quand l'utilisteur click sur 'enter'
		if err != nil {
			return choix, err
		}
		choix = strings.TrimSpace(choix)
		if len(choix) != 1 {
			fmt.Printf("saisisez un nombre valide de caractère \n carac choisi : %v \nlongeur de carac : %v ", choix, len(choix))
			continue
		}
		valid = true

	}

	return choix, nil // nil parceque toute les erreurs ont été traité en amont
}
