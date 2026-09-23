package player

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/common"
	"strings"
	"unicode"
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
		a.MONNAIE = 100
	case 1:
		a.TYPE = "chartaceus"
		a.PVMAX = 90
		a.PV = a.PVMAX / 2
		a.ATK = 12
		a.DEF = 1
		a.PRMAX = 150
		a.PR = 0
		a.MONNAIE = 100
	case 2:
		a.TYPE = "plantyrien"
		a.PVMAX = 120
		a.PV = a.PVMAX / 2
		a.ATK = 8
		a.DEF = 3
		a.PRMAX = 50
		a.PR = 0
		a.MONNAIE = 100
	default:
		fmt.Println("\nErreur de saisi pour le Type de personnage, veuillez saisir le numéro du Type de personnage ou sont nom complet sans erreur de frappe")
	}
}

// définition des variable qui possèderont les infos a saisir
var TypeSaisi int = 0
var NomSaisi string
var test Character

func EnterName() string {
	NomSaisi = ""
	for {
		fmt.Println("\nVeuillez Saisir votre nom:")
		NomSaisi = bubble.StartInput("\tJe m'appelle....\t")
		NomSaisi = strings.TrimSpace(NomSaisi)
		if VerifName(NomSaisi) {
			runes := []rune(NomSaisi)
			ValideNom := string(runes[0]) + strings.ToLower(string(runes[1:]))
			NomSaisivalide := strings.Title(ValideNom)
			return NomSaisivalide
		}
	}
}

func VerifName(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		fmt.Println("\nErreur : Le nom ne peut pas être vide.")
		return false
	}
	// Bloque les caractères spéciaux, chiffres et espaces (seules les lettres sont autorisées)
	for _, a := range runes {
		if !unicode.IsLetter(a) {
			fmt.Println("\nErreur de Saisi. Le nom ne doit pas contenir de caractères spéciaux ou de chiffres.")
			return false
		}
	}
	return true
}

func EnterType() int {
	common.DisplayNarration(("Bienvenue à bord de votre vaisseau " + NomSaisi + ". Vous êtes entré(e) dans un monde où nul à de logique."))

	TypeSaisi := bubble.StartChoice([]string{"Humain.[Personnage par défaut, rien de spécial.]",
		"Chartaceus.[Un Homme-papier, faible mais coupe plus facilment.]",
		"Plantyrien.[Une plante qui aspire les nutriments d'autrui. Peut survivre plus longtemps.]"}, false)

	return TypeSaisi
}
