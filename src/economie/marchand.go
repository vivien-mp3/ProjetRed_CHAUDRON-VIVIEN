package economie

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/item"
	"projet/src/player"
)

var stockJus = 2
var stockPizza = 2

// Func permetant la création d'une boutique ( Marchand )
func Boutique(p *player.Character) {
	dansBoutique := true
	for dansBoutique {
		fmt.Println("Bienvenue jeune Inconnu, voici ma boutique. Que voulez-vous acheter ?")
		fmt.Println("Joueur :", p.NAME, "| Monnaie :", p.MONNAIE)
		//permet d'afficher le stock de jus d'ananas
		texteJus := fmt.Sprintf("Jus d'ananas (x%d) - 15 écus", stockJus)
		if stockJus == 0 {
			texteJus = "Jus d'ananas [ÉPUISÉ]"
		}
		//permet d'afficher le stock de la pizza d'ananas
		textePizza := fmt.Sprintf("Pizza à l'ananas (x%d) - 15 écus", stockPizza)
		if stockPizza == 0 {
			textePizza = "Pizza à l'ananas [ÉPUISÉ]"
		}

		options := []string{texteJus, textePizza, "Quitter la boutique"}
		choix := bubble.StartChoice(options, false)
		// permet la possibilité d'interagire et payer avec la monnaie 
		switch choix {
		case 0:
			if stockJus > 0 {
			if p.MONNAIE >= item.Jus.Prix {
				RetraitMonnaie(p, item.Jus.Prix)
				player.AddInventory(item.Jus.Nom)
					stockJus--
					fmt.Println("Vous avez acheté un", item.Jus.Nom, "(Reste en stock :", stockJus, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Jus d'ananas.")
			}

		case 1:
			if stockPizza > 0 {
			if p.MONNAIE >= item.Pizza.Prix {
				RetraitMonnaie(p, item.Pizza.Prix)
				player.AddInventory(item.Pizza.Nom)
					stockPizza--
					fmt.Println("Vous avez acheté une", item.Pizza.Nom, "(Reste en stock :", stockPizza, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Pizza à l'ananas.")
			}

			case 2:
				fmt.Println("Merci de votre visite !")
				dansBoutique = false
			}
		}
}
