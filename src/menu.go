package src

//Code permetant la création d'un menu start, qui permet de choisir le personnage et de lancer le jeu
import (
	"fmt"
	"os"
	"projet/src/bubble"
	"projet/src/common"
)

func Startmenu() {
	common.PlaySound("deepspace")
	//code permetant d'arriver sur le jeu .
	var choice int

	for result := true; result; {
		fmt.Println("\033[1;35m")
		fmt.Println(`
	 ____             _      _     ____    ___               
	|  _ \ _ __ ___  (_) ___| |_  |  _ \  /_/_\   _____  
	| |_) | '__/ _ \ | |/ _ \ __| | |_) |/ _ \ \ / / _ \  
	|  __/| | | (_) || |  __/ |_  |  _ <|  __/\ V /  __/  
	|_|   |_|  \___/_/ |\___|\__| |_| \_\\___| \_/ \___|  
	               |__/                                    
	`)
		fmt.Println("\033[0m")
		fmt.Println("Bienvenue dans le jeu!")
		options := []string{"Commencer une nouvelle partie", "Parametre", "Quitter le jeu"}
		choice = bubble.StartChoice(options, false)

		switch choice {

		case 0:
			fmt.Println("Nouvelle partie commencée !")
			result = false
		case 1:
			fmt.Println("Parametre")
		case 2:
			fmt.Println("Vous avez quittez le jeu")
			result = false
			os.Exit(0)
		default:
			fmt.Println("Oups tromper de touche")
		}
	}
}
