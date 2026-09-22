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

	
	options := []string{"Gauche", "Droite"}

	choix := bubble.StartChoice(options, false)
	switch choix {
	case 0 :
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à gauche.")
			fmt.Println("Vous appercevez une lumiere , un sorte village.")
			fmt.Println("\033[0m")
		
		case 1:
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à droite.")
			fmt.Println("\033[0m")
	}

}

func Cartegauche(){

	fmt.Println("Arriver au village vous croisser un garde")
	fmt.Println("Garde : Alte inconnue , que vaut votre présence dans ce village")
	fmt.Println("[*name] Je recherche un endroit ou passer la nuit")
	fmt.Println("Garde : Vous avez une auberge au fond, ils vont vous accueillir !")

	options := []string{"Rentrez dans l'Auberge", "Partir du village"}
	choix := bubble.StartChoice(options, false)

	switch choix {
	case 0 :
		fmt.Println("\033[1;35")
		fmt.Println("Vous rentrez dans l'auberge est vous appercevez un Barman")
		fmt.Println("\033[0m")

		case 1 :
			fmt.Println("\033[1;35m")
			fmt.Println("Quittez le village et explorer les environ malgrés les environ obscur")
			fmt.Println("\033[0m")
	}
}