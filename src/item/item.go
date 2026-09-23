package item

import ()

type  Item struct{
	Nom string
	Prix int
	Description string
}

var Jus = Item{
	Nom : "Jus d'ananas",
	Prix: 15, 
	Description:"Permet de rajouter + 20 PV",
}

var Pizza = Item{
	Nom: "Pizza à l'ananas",
	Prix: 15,
	Description: "retire -25 PV à son adversaire",
}