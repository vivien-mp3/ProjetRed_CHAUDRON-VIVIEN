package src

import "fmt"

var icons = map[string]string{"chat": "ᓚᘏᗢ", "joueur": "𐀪"}

func DisplayDialogue(icon string, m string) {
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]\n")
	fmt.Printf("|[ %s ]|\n\t<< %s >>\n", icons[icon], m)
}

func DisplayNarration(m string){
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]\n")
	fmt.Printf("[[ %s ]]\n", m)
}

