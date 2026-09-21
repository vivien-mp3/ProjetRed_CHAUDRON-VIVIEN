package src

//Code permetant la création d'un menu start, qui permet de choisir le personnage et de lancer le jeu

import (
	"fmt"
	"os"
)

<<<<<<< HEAD
	import (
		"fmt"
		"os"
		
	)

	func StartMenu(){
		//code permetant d'arriver sur le jeu .
		var choice int
		for result := true ; result; {
=======
func StartMenu() {
	//code permetant d'arriver sur le jeu .
	var choice int
	for result := true; result; {
>>>>>>> 25be4df63352b9fcf3405a1bd6a98c628f143734
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
<<<<<<< HEAD
		default :
		fmt.Println("Oups tromper de touche")
		}		
	}	
}
=======
		default:
			fmt.Println("Oups tromper de touche\n")
		}
	}
}
>>>>>>> 25be4df63352b9fcf3405a1bd6a98c628f143734
