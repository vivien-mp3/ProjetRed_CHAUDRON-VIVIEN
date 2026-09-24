package item

type Item struct {
	Nom         string
	Prix        int // Monnaie item
	Description string
	Soin        int // PV rendus au joueur
	Degats      int // Dégâts infligés à l'adversaire
}

// Variable permettant de créer le Jus d'ananas (+20 PV)
var Jus = Item{
	Nom:         "Jus d'ananas",
	Prix:        15,
	Description: "Permet de rajouter + 20 PV",
	Soin:        20,
	Degats:      0,
}

// Variable permettant de créer la Pizza à l'ananas (-25 PV adversaire)
var Pizza = Item{
	Nom:         "Pizza à l'ananas",
	Prix:        15,
	Description: "retire -25 PV à son adversaire",
	Soin:        0,
	Degats:      25,
}

// Variable permettant de créer le Sac à Dos (+10 taille inventaire)
var Sac = Item{
	Nom:         "Sac à Dos",
	Prix:        30,
	Description: "Augmente de 5 la taille de l'inventaire",
}
