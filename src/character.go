package src

/*
La commande pour afficher les stats du personnage est la suivante
var test src.Character
	test.InitCharacter("NOOB", "1")
	fmt.Println(test)
	vous pouvez remplacer le "test" par le nom de la variable de votre choix
*/

import (
	"fmt"
	"strings"
)

// struct avec tous les caractères des personnages
type Character struct {
	NAME  string
	TYPE  string
	PVMAX int
	PV    int
	ATK   int
	DEF   int
	PRMAX int
	PR    int
	INV   map[string]int
}

// défini la variable Bug pour le cas ou le personnage sélectionné n'existe pas
var Bug bool

// défini les stats du personnage correspondant au personnage choisi, choix possible avec le nom ou le numéro du personnage
func (a *Character) InitCharacter(name string, types string) {
	a.NAME = name
	types = strings.ToLower(types)
	if types == "1" || types == "humain" {
		a.TYPE = "humain"
		a.PVMAX = 100
		a.PV = a.PVMAX / 2
		a.ATK = 10
		a.DEF = 2
		a.PRMAX = 100
		a.PR = 0
	} else if types == "2" || types == "chartaceus" {
		a.TYPE = "chartaceus"
		a.PVMAX = 90
		a.PV = a.PVMAX / 2
		a.ATK = 12
		a.DEF = 1
		a.PRMAX = 150
		a.PR = 0
	} else if types == "3" || types == "plantyrien" {
		a.TYPE = "plantyrien"
		a.PVMAX = 120
		a.PV = a.PVMAX / 2
		a.ATK = 8
		a.DEF = 3
		a.PRMAX = 50
		a.PR = 0
	} else {
		fmt.Println("Erreur de saisi pour le Type de personnage, veuillez saisir le numéro du Type de personnage ou sont nom complet sans erreur de frappe")
		Bug = true
	}
}
