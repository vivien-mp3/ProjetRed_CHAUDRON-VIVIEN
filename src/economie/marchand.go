package economie

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/item"
	"projet/src/player"
)

var stockSacADos = 3

// Func permetant la création d'une boutique ( Marchand )
func Boutique(p *player.Character) {
	var stockJus = 2
	var stockPizza = 2
	var stockGriffe = 2
	var stockMoteur = 2
	var stockSoie = 2
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
		//permet d'afficher le stock de sac à dos
		texteSac := fmt.Sprintf("Sac à Dos (x%d) - 30 écus", stockSacADos)
		if stockSacADos == 0 {
			texteSac = "Sac à Dos [ÉPUISÉ]"
		}
		//permet d'afficher le stock de moteur
		texteMoteur := fmt.Sprintf("Moteur (x%d) - 10 écus", stockMoteur)
		if stockMoteur == 0 {
			texteMoteur = "Moteur [ÉPUISÉ]"
		}
		//permet d'afficher le stock de griffe de griffon
		texteGriffe := fmt.Sprintf("Griffe de Griffon (x%d) - 10 écus", stockGriffe)
		if stockGriffe == 0 {
			texteGriffe = "Griffe de griffon [ÉPUISÉ]"
		}
		//permet d'afficher le stock de soie
		texteSoie := fmt.Sprintf("Soie de Bombyx (x%d) - 10 écus", stockSoie)
		if stockSoie == 0 {
			texteSoie = "Soie de Bombyx [ÉPUISÉ]"
		}

		options := []string{texteJus, textePizza, texteSac, texteMoteur, texteGriffe, texteSoie, "Quitter la boutique"}
		choix := bubble.StartChoice(options, false)
		// permet la possibilité d'interagire et payer avec la monnaie
		switch choix {
		case 0:
			if stockJus > 0 {
				if p.MONNAIE >= item.JusDeBanane.Price {
					RetraitMonnaie(p, item.JusDeBanane.Price)
					player.AddInventory(item.JusDeBanane.Name, 1, &p.INV)
					stockJus--
					fmt.Println("Vous avez acheté un", item.JusDeBanane.Name, "(Reste en stock :", stockJus, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Jus d'ananas.")
			}

		case 1:
			if stockPizza > 0 {
				if p.MONNAIE >= item.Pizza.Price {
					RetraitMonnaie(p, item.Pizza.Price)
					player.AddInventory(item.Pizza.Name, 1, &p.INV)
					stockPizza--
					fmt.Println("Vous avez acheté une", item.Pizza.Name, "(Reste en stock :", stockPizza, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Pizza à l'ananas.")
			}

		case 2:
			if stockSacADos > 0 {
				if p.MONNAIE >= item.Sac.Price {
					RetraitMonnaie(p, item.Sac.Price)
					player.UpTailleInv()
					stockSacADos--
					fmt.Println("Vous avez acheté un", item.Sac.Name, "(Reste en stock :", stockSacADos, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Sac à Dos.")
			}

		case 3:
			if stockMoteur > 0 {
				if p.MONNAIE >= item.Moteur.Price {
					RetraitMonnaie(p, item.Moteur.Price)
					player.AddInventory(item.Moteur.Name, 1, &p.INV)
					stockMoteur--
					fmt.Println("Vous avez acheté un", item.Moteur.Name, "(Reste en stock :", stockMoteur, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Moteur.")
			}

		case 4:
			if stockGriffe > 0 {
				if p.MONNAIE >= item.Griffe.Price {
					RetraitMonnaie(p, item.Griffe.Price)
					player.AddInventory(item.Griffe.Name, 1, &p.INV)
					stockGriffe--
					fmt.Println("Vous avez acheté un", item.Griffe.Name, "(Reste en stock :", stockGriffe, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Griffe de griffon.")
			}

		case 5:
			if stockSoie > 0 {
				if p.MONNAIE >= item.Soie.Price {
					RetraitMonnaie(p, item.Soie.Price)
					player.AddInventory(item.Soie.Name, 1, &p.INV)
					stockSoie--
					fmt.Println("Vous avez acheté un", item.Soie.Name, "(Reste en stock :", stockSoie, ")")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent !")
				}
			} else {
				fmt.Println("Rupture de stock ! Il n'y a plus de Soie de Bombyx.")
			}

		case 6:
			fmt.Println("Merci de votre visite !")
			dansBoutique = false
		}
	}
}
