package player

import (
	"fmt"
	"projet/src/bubble"

	//"projet/src/item"
	"strconv"
	"strings"
)

//strconv permet de transformer le int de value en string afin de pouvoir l'afficher

var ListInv []string
var ItemPro string

// var INV = make(map[string]int)
var temp []string

func AccesInventory(INV *[]Inventory) {
	for true {
		ListInv = nil
		temp = nil
		for _, it := range *INV {
			ItemPro = it.Name + " x" + strconv.Itoa(it.Quantity)
			temp = append(temp, it.Name)
			ListInv = append(ListInv, ItemPro)
		}
		ListInv = append(ListInv, "Retour.")
		selec := bubble.StartChoice(ListInv, false)
		if ListInv[selec] == "Retour." {
			fmt.Println("Vous avez fermé votre inventaire.")
			return
		}
		for a := 0; a < len(ListInv); a++ {
			if selec == a {
				objet := temp[a]
				SupInventory(objet, 1, INV)
			}
		}
	}
}

// Ceci est la variable avec le nombre d'item dans l'inventaire
var LIMITINV int = 10

func AddInventory(obj string, objqt int, INV *[]Inventory) {
	isItemHere := false
	nbItem := 0
	for i, item := range *INV {
		nbItem += item.Quantity
		if strings.EqualFold(obj, item.Name) {
			isItemHere = true
			jaja := *INV //JE DETESTE LE GOLANG
			jaja[i].Quantity += objqt
			fmt.Printf("+ %s x %d a été rajouté avec vos autres objets inventaire.\n", obj, objqt)
		}
		fmt.Println(item.Name, item.Quantity, item)
	}
	if nbItem > LIMITINV {
		fmt.Println("Vous venez d'atteindre le maximum de votre inventaire")
		return
	}
	if isItemHere {
		return
	}
	*INV = append(*INV, Inventory{obj, objqt})
	fmt.Printf("%s x %d a été ajouté à votre inventaire.\n", obj, objqt)
}

func SupInventory(obj string, objqt int, INV *[]Inventory) {
	var tempoINV []Inventory
	for _, item := range *INV {
		if strings.EqualFold(obj, item.Name) {
			if item.Quantity > objqt {
				item.Quantity -= objqt
				tempoINV = append(tempoINV, item)
			}
		} else {
			tempoINV = append(tempoINV, item)
		}
	}
	*INV = tempoINV
}

func UpTailleInv() {
	LIMITINV += 5
}
