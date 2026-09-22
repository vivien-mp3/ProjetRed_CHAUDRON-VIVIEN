package src

import (
	"fmt"
	"projet/src/bubble"
)

type CarteStruct struct {
	NAME        string
	Description string
	Evenement   int
	Lieux       bool
}

func Carte() {
	// 1. On définit les options proposées au joueur
	options := []string{"Gauche", "Droite"}

	// 2. On lance le menu interactif Bubble Tea (flèches + Entrée)
	// Le 2e argument (false) indique si on veut afficher l'option inventaire ou non
	choix := bubble.StartChoice(options, false)

	// 3. On traite le choix selon l'index retourné (0 = Gauche, 1 = Droite)
	switch choix {
	case 0:
		fmt.Println("Vous avez choisi d'aller à gauche.")
	case 1:
		fmt.Println("Vous avez choisi d'aller à droite.")
	}
}
