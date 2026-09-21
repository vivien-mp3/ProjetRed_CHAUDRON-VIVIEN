package src

import (
	"fmt"
	"strings"
	"time"
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
	if types == "0" || types == "humain" {
		a.TYPE = "humain"
		a.PVMAX = 100
		a.PV = a.PVMAX / 2
		a.ATK = 10
		a.DEF = 2
		a.PRMAX = 100
		a.PR = 0
	} else if types == "1" || types == "chartaceus" {
		a.TYPE = "chartaceus"
		a.PVMAX = 90
		a.PV = a.PVMAX / 2
		a.ATK = 12
		a.DEF = 1
		a.PRMAX = 150
		a.PR = 0
	} else if types == "2" || types == "plantyrien" {
		a.TYPE = "plantyrien"
		a.PVMAX = 120
		a.PV = a.PVMAX / 2
		a.ATK = 8
		a.DEF = 3
		a.PRMAX = 50
		a.PR = 0
	} else {
		fmt.Println("\nErreur de saisi pour le Type de personnage, veuillez saisir le numéro du Type de personnage ou sont nom complet sans erreur de frappe")
		Bug = true
	}
}

/*
J'utilise la commande fmt.Scanln() afin de saisir les infos nécessaire pour le personnage
*/

// définition des variable qui possèderont les infos a saisir
var TypeSaisi string = "debug"
var NomSaisi string = "debug"
var test Character

func SaisiInfo() {
	for {
		// choix du type de personnage
		fmt.Println("\nVeuillez Saisir le type de personnage entre les différents proposé :")
		fmt.Println("1 - Humain || 2 - chartaceus || 3 - plantyrien")
		fmt.Scanln(&TypeSaisi)
		test.InitCharacter(NomSaisi, TypeSaisi)
		if Bug == true {
			Bug = false
			time.Sleep(2 * time.Second)
			continue
		} else {
			break
		}
	}
	//choix du nom du personnage
	time.Sleep(1 * time.Second)
	fmt.Println("\nVeuillez Saisir le nom de votre personnage entre les différents proposé :")
	fmt.Scanln(&NomSaisi)
	fmt.Println("\nLe nom de votre personnage est : ", NomSaisi, "\n")
}
