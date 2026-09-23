package src

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/player"
)

type CarteStruct struct {
	NAME        string
	Description string
	Evenement   int
	Lieux       bool
}
// création de la fonction Carte permettant les déplacement de ce personnage , utilisant bubbletea
// p *player.Character permet d'appeler le ficher Character afin de m'afficher le nom du personnage dans les discusion 
func Carte(p *player.Character) {
	options := []string{"Gauche", "Droite"}

	choix := bubble.StartChoice(options, false)
	// Permet la selection de choix entre tourner à droite ou à gauche
	switch choix {
	case 0 :
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à gauche\n.")
			fmt.Println("\033[0m")
			fmt.Println("\033[1;31m")
			fmt.Println(p.NAME,": Vous appercevez une lumiere , un sorte village.")
			fmt.Println("\033[0m")
			Cartegauche(p)
		
		case 1:
			fmt.Println("\033[1;35m")
			fmt.Println("Vous avez choisi d'aller à droite.")
			fmt.Println("\033[0m")
	}

}

func Cartegauche(p *player.Character) {

	
	fmt.Println("Arriver au village vous croisser un garde")

	fmt.Println("\033[1;32m")
	fmt.Println("Garde : Alte inconnue , que vaut votre présence dans ce village")
	fmt.Println("\033[0m")

	fmt.Println("\033[1;31m")
	fmt.Println(p.NAME,": Je recherche un endroit ou passer la nuit")
	fmt.Println("\033[0m")
	
	fmt.Println("\033[1;32m")
	fmt.Println("Garde : Vous avez une auberge au fond, ils vont vous accueillir !")
	fmt.Println("\033[0m")

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