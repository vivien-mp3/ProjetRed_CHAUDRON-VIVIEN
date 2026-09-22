package src
import (
	"fmt" 
	"os"
)
	type carte struct{
	NAME  string
	Description  string
	Evenement  int
	Lieux bool
	}
func Carte(){
	var Chemin int

	fmt.Println("Quel chemin allez-vous emprunter ?\n")
	fmt.Println("1 : Gauche ")
	fmt.Println("2 : Doitre")
	fmt.Scanln(&Chemin)

	for result := true; result; {
	switch Chemin {
		case 1:
			fmt.Println("Vous avez choisie de tourner  gauche \n")
			fmt.Println("NAME : TEST \n")
			fmt.Println("Description : Vous arriver dans un endroit sombre et mysterieux\n ")
			result = false

		case 2 :
			fmt.Println("Vous avez choisie de tourner à droite \n")

		default:
			fmt.Println("La direction sellectioner est inconnue")
			os.Exit(0)
		}
	}
}