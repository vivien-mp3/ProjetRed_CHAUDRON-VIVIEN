package player

import (
	"projet/src/bubble"
	"strconv"
)

//strconv permet de transformer le int de value en string afin de pouvoir l'afficher

var ListInv []string
var ItemPro string
var INV = make(map[string]int)

func AccesInventory() {
	for key, value := range INV {
		ItemPro = key + " x" + strconv.Itoa(value)
		ListInv = append(ListInv, ItemPro)
	}
	selec := bubble.StartChoice(ListInv, false)
	ListInv = nil
	switch selec {
		
	}
}

func AddInventory(objet string) {
	INV[objet]++
}

func SupInventory(objet string) {
	INV[objet]--
	if INV[objet] == 0 {
		delete(INV, objet)
	}
}
