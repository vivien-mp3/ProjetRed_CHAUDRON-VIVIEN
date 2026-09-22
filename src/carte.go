package src

import (
	"fmt"
	"projet/src/bubble"
)

type CarteStruct struct {
	NAME        string
	Description string
	Evenement   int
	Lieux       bool
}

func Carte() {
	var position int

	
	options := []string{"Gauche", "Droite"}

	choix := bubble.StartChoice(options, false)
	switch choix {
	case 0 :
		if position > 0 {
			position--
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à gauche.")
			fmt.Println("Vous appercevez une lumiere , un sorte village.")
			fmt.Println("\033[0m")
		}
		case 1:
		if position < -1{
				position++
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à droite.")
			fmt.Println("\033[0m")
		}
	}
}
