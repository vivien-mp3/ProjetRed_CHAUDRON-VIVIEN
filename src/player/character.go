package player

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/common"
	"strings"
)

// struct avec tous les caractères des personnages
type Character struct {
	NAME    string
	TYPE    string
	PVMAX   int
	PV      int
	ATK     int
	DEF     int
	PRMAX   int
	PR      int
	INV     map[string]int
	MONNAIE int
}

// défini les stats du personnage correspondant au personnage choisi, choix possible avec le nom ou le numéro du personnage
func (a *Character) InitCharacter(name string, types int) {
	a.NAME = name
	switch types {
	case 0:
		a.TYPE = "humain"
		a.PVMAX = 100
		a.PV = a.PVMAX / 2
		a.ATK = 10
		a.DEF = 2
		a.PRMAX = 100
		a.PR = 0
	case 1:
		a.TYPE = "chartaceus"
		a.PVMAX = 90
		a.PV = a.PVMAX / 2
		a.ATK = 12
		a.DEF = 1
		a.PRMAX = 150
		a.PR = 0
	case 2:
		a.TYPE = "plantyrien"
		a.PVMAX = 120
		a.PV = a.PVMAX / 2
		a.ATK = 8
		a.DEF = 3
		a.PRMAX = 50
		a.PR = 0
	default:
		fmt.Println("\nErreur de saisi pour le Type de personnage, veuillez saisir le numéro du Type de personnage ou sont nom complet sans erreur de frappe")
	}
}

/*
J'utilise la commande fmt.Scanln() afin de saisir les infos nécessaire pour le personnage
*/

// définition des variable qui possèderont les infos a saisir
var TypeSaisi int = 0
var NomSaisi string
var test Character
var NomSaisin string

func EnterName() string {
	for true {

		fmt.Println("\nVeuillez Saisir votre nom:")
		NomSaisin = bubble.StartInput("\tJe m'appelle....\t")
		NomSaisin = strings.ToLower(NomSaisin)
		for abcdef := 0; abcdef < len(NomSaisin); abcdef++ {
			if abcdef == 0 {
				NomSaisi += string(rune(NomSaisin[0] - 32))
			} else {
				NomSaisi += string(NomSaisin[abcdef])
			}
		}
		fmt.Println("\nVotre nom est donc: ", NomSaisi)
		return NomSaisi
	}
	return NomSaisi
}

func EnterType() int {
	common.DisplayNarration(("Bienvenue à bord de votre vaisseau " + NomSaisi + ". Vous êtes entré(e) dans un monde où nul à de logique."))

	TypeSaisi := bubble.StartChoice([]string{"Humain.[Personnage par défaut, rien de spécial.]",
		"Chartaceus.[Un Homme-papier, faible mais coupe plus facilment.]",
		"Plantyrien.[Une plante qui aspire les nutriments d'autrui. Peut survivre plus longtemps.]"}, false)

	return TypeSaisi
}
