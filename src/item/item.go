package item

import (
	"fmt"
	"projet/src/player"
)

type Item struct {
	Name  string
	Price int // Monnaie item
	Heal  int // PV rendus au joueur
	Dmg   int // Dégâts infligés à l'adversaire
}

var JusDeBanane = Item{
	Name:  "Jus de Banane",
	Price: 15,
	Heal:  20,
	Dmg:   0,
}

// Variable permettant de créer le Sac à Dos (+10 taille inventaire)
var Sac = Item{
	Name:  "Sac à Dos",
	Price: 30,
}

var Moteur = Item{
	Name:  "Moteur",
	Price: 10,
}

var Griffe = Item{
	Name:  "Griffe de Griffon",
	Price: 10,
}

var Soie = Item{
	Name:  "Soie de Bombyx",
	Price: 10,
}

func UseItem(obj string, INV *[]player.Inventory, plr *player.Character) {
	for _, it := range *INV {
		fmt.Println(obj, it)
		if it.Name == obj {
			if it.Quantity > 0 {
				it.Quantity -= 1
				switch it.Name {
				case "Jus de banane":
					plr.PV += JusDeBanane.Heal
					if plr.PV > plr.PVMAX {
						plr.PV = plr.PVMAX
					}
				case "Pizza a l'ananas":
					
				default:
					fmt.Printf("|%s ne peut pas être utilisé.\n", it.Name)
					return
				}
			}
		}
	}
}

//Fonction useHeal inutilisable.
/* func useHeal(heal int, plr player.Character) int {
	newHP := plr.PV + heal
	if newHP > plr.PVMAX {
		newHP = plr.PVMAX
	}
	return newHP
} */
