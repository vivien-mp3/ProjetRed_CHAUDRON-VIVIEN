package item

import (
	"fmt"
	"projet/src/bubble"

	//"projet/src/item"
	"projet/src/player"
)

var stockHelmet = 1
var stockChesplate = 1
var stockBoots = 1

// fonction pour vérifier la possession des items requis
func AItem(itemName string) bool {
	for _, item := range player.Statplr.INV {
		if item.Name == itemName && item.Quantity > 0 {
			return true
		}
	}
	return false
}

// fonction de trade pour le craft d'item
func Mecano() {
	dansTrain := true
	for dansTrain {

		texteHelmet := fmt.Sprintf("Helmet (x%d) - 1 Griffe de Griffon & 1 Soie de Bombyx", stockHelmet)
		if stockHelmet == 0 {
			texteHelmet = "Helmet [ÉPUISÉ]"
		}

		texteChesplate := fmt.Sprintf("Chesplate (x%d) - 1 Moteur & 1 Soie de Bombyx", stockChesplate)
		if stockChesplate == 0 {
			texteChesplate = "Chesplate [ÉPUISÉ]"
		}

		texteBoots := fmt.Sprintf("Boots (x%d) - 1 Moteur & 1 Griffe de Griffon", stockBoots)
		if stockBoots == 0 {
			texteBoots = "Boots [ÉPUISÉ]"
		}

		options := []string{texteHelmet, texteChesplate, texteBoots, "Quitter la boutique"}
		choix := bubble.StartChoice(options, false)
		switch choix {
		case 0:
			if stockHelmet > 0 {
				if AItem(Soie.Name) && AItem(Griffe.Name) {
					player.AddInventory(Helmet.Name, 1, &[]player.Inventory{})
					player.SupInventory(Soie.Name, 1, &[]player.Inventory{})
					player.SupInventory(Griffe.Name, 1, &[]player.Inventory{})
					stockHelmet--
					fmt.Println("Vous avez acheté un", Helmet.Name, "(Reste en stock :", stockHelmet, ")")
				} else {
					fmt.Println("Vous n'avez pas les items requis !")
					continue
				}
			}

		case 1:
			if stockChesplate > 0 {
				if AItem(Moteur.Name) && AItem(Soie.Name) {
					player.AddInventory(Chesplate.Name, 1, &[]player.Inventory{})
					player.SupInventory(Moteur.Name, 1, &[]player.Inventory{})
					player.SupInventory(Soie.Name, 1, &[]player.Inventory{})
					stockChesplate--
					fmt.Println("Vous avez acheté un", Chesplate.Name, "(Reste en stock :", stockChesplate, ")")
				} else {
					fmt.Println("Vous n'avez pas les items requis !")
					continue
				}
			}

		case 2:
			if stockBoots > 0 {
				if AItem(Moteur.Name) && AItem(Griffe.Name) {
					player.AddInventory(Boots.Name, 1, &[]player.Inventory{})
					player.SupInventory(Moteur.Name, 1, &[]player.Inventory{})
					player.SupInventory(Griffe.Name, 1, &[]player.Inventory{})
					stockBoots--
					fmt.Println("Vous avez acheté un", Boots.Name, "(Reste en stock :", stockBoots, ")")
				} else {
					fmt.Println("Vous n'avez pas les items requis !")
					continue
				}
			}

		case 3:
			fmt.Println("Merci de votre visite !")
			dansTrain = false
		}
	}
}
