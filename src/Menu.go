package src

//Code permetant la création d'un menu start, qui permet de choisir le personnage et de lancer le jeu
import (
	"fmt"
	"os"
)

func StartMenu() {
	//code permetant d'arriver sur le jeu .
	var choice int

	for result := true; result; {
		fmt.Println("\033[1;35m")
		fmt.Println("╔══════════════════════════════╗")
		fmt.Println("║         Projet Rêve          ║")
		fmt.Println("╚══════════════════════════════╝")
		fmt.Println("\033[0m")
		fmt.Println("Bienvenue dans le jeu!")
		fmt.Println("1 : Commencer une nouvelle partie")
		fmt.Println("2 : Charger une partie")
		fmt.Println("3 : Quitter le jeu")
		fmt.Print("Veuillez entrer votre choix (1, 2 ou 3) : ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			fmt.Println("Nouvelle partie commencée !")
			result = false
		case 2:
			fmt.Println("Parametre")
		case 3:
			fmt.Println("Vous avez quittez le jeu")
			result = false
			os.Exit(0)
		default:
			fmt.Println("Oups tromper de touche\n")
		}
	}
}
