package economie

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/item"
	"projet/src/player"
)

func Boutique(p *player.Character) {
	dansBoutique := true
	for dansBoutique {
		fmt.Println("Bienvenue jeune Inconnu, voici ma boutique. Que voulez-vous acheter ?")
		fmt.Println("Joueur :", p.NAME, "| Monnaie :", p.MONNAIE)

		options := []string{"Jus d'ananas - 15 écus", "Pizza à l'ananas - 15 écus", "Quitter la boutique"}
		choix := bubble.StartChoice(options, false)

		switch choix {
		case 0:
			if p.MONNAIE >= item.Jus.Prix {
				RetraitMonnaie(p, item.Jus.Prix)
				player.AddInventory(item.Jus.Nom)
				fmt.Println("Vous avez acheté un", item.Jus.Nom)
			} else {
				fmt.Println("Vous n'avez pas assez d'argent !")
			}
		case 1:
			if p.MONNAIE >= item.Pizza.Prix {
				RetraitMonnaie(p, item.Pizza.Prix)
				player.AddInventory(item.Pizza.Nom)
				fmt.Println("Vous avez acheté une", item.Pizza.Nom)
			} else {
				fmt.Println("Vous n'avez pas assez d'argent !")
			}
		case 2:
			fmt.Println("Merci de votre visite !")
			dansBoutique = false
		}
	}

}
