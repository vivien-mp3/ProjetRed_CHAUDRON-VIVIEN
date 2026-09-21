package src

import "fmt"


// Ici sont placé les icones des personnages qui parle
var icons = map[string]string{"chat": "ᓚᘏᗢ", "joueur": "𐀪"}


// Ici est la fonction pour les dialogues, icon = personnage qui parle, t = le nom du personnage qui parle, m = le message
func DisplayDialogue(icon string, t string, m string) {
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]\n")
	fmt.Printf("|[ %s ]| -- %s o--> \n\t<< %s >>\n", icons[icon], t, m)
}

// Ici est la fonction pour les dialogues
func DisplayNarration(m string){
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]\n")
	fmt.Printf("\t[[ %s ]]\n", m)
}

