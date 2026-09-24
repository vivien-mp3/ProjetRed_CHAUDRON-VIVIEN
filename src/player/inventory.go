package player

import (
	"fmt"
	"projet/src/bubble"
	"strconv"
)

//strconv permet de transformer le int de value en string afin de pouvoir l'afficher

var ListInv []string
var ItemPro string
var INV = make(map[string]int)
var temp []string

func AccesInventory() {
	for true {
		for key, value := range INV {
			ItemPro = key + " x" + strconv.Itoa(value)
			temp = append(temp, key)
			ListInv = append(ListInv, ItemPro)
		}
		ListInv = append(ListInv, "retour")
		selec := bubble.StartChoice(ListInv, false)
		if ListInv[selec] == "retour" {
			fmt.Println("Vous avez fermé votre inventaire.")
			return
		}
		for a := 0; a < len(ListInv); a++ {
			if selec == a {
				objet := temp[a]
				SupInventory(objet)
			}
		}
		ListInv = nil
		temp = nil
	}
}

// Ceci est la variable avec le nombre d'item dans l'inventaire
var NbItemInv int = 0
var LIMITINV int = 10

func AddInventory(objet string) {
	if NbItemInv < LIMITINV {
		INV[objet]++
		NbItemInv += 1
		if NbItemInv == LIMITINV {
			fmt.Println("Vous venez d'atteindre le maximum de votre inventaire")
		}
	} else {
		fmt.Println("Limite d'inventaire atteint vous ne pouvez pas ajouter d'item")
	}
}

func SupInventory(objet string) {
	INV[objet]--
	if INV[objet] == 0 {
		delete(INV, objet)
	}
	NbItemInv -= 1
}

func UpTailleInv() {
	LIMITINV += 5
}
