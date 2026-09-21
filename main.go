package main

import (
	"projet/src"
	"projet/src/common"
)

type character struct {
	pc int
}

func main() {
	//permet initialiser le menu
	//src.StartMenu()
	// permet de lancer la narration
	common.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	common.DisplayDialogue("joueur", "Le joueur", "Oe oe oe !")
	mess := "Un vent froid coule sur votre peau.\n\tVous vous sentez bizarre, le front chaud, mais le corps froid.\n\tCela semble comme un cauchemard."
	common.DisplayNarration(mess)
	go common.PlaySound("deepspace")
	for true {
		answer := src.StartChoice([]string{"Test1", "Test2"}, true)
		switch answer{
		case 0: common.DisplayNarration("Vous avez choisi le test numéro 1.")
		common.StopAllSounds()
		case 1: common.DisplayDialogue("chat", "le chat", "Meow meow meow. Le choix numéro 2 !")
		case 2: common.DisplayNarration("Vous accédez à l'inventaire.")
	}
	
	}
}
