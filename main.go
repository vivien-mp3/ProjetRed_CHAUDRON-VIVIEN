package main

import (
	"fmt"
	"projet/src"
	chrfol "projet/src/character"
	"projet/src/common"
	"time"
)

type character struct {
	pc int
}

func main() {
	//permet initialiser le menu
	//src.StartMenu()
	//requete pour le personnage
	var Test *string
	fmt.Scanln(&Test)
	plrName := chrfol.EnterName()
	plrType := chrfol.EnterType()
	//information sur le personnage
	var plr chrfol.Character
	plr.InitCharacter(plrName, plrType)
	// permet de lancer la narration
	common.DisplayInfo(plr.NAME, plr.TYPE, plr.PVMAX, plr.PV, plr.ATK, plr.DEF)
	time.Sleep(1 * time.Second)
	common.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	time.Sleep(2 * time.Second)
	common.DisplayDialogue("joueur", chrfol.NomSaisi, "Oestory. oe oe !")
	time.Sleep(1 * time.Second)
	mess := "Un vent froid coule sur votre peau.\n\tVous vous sentez bizarre, le front chaud, mais le corps froid.\n\tCela semble comme un cauchemard."
	common.DisplayNarration(mess)
	go common.PlaySound("deepspace")
	for true {
		answer := src.StartChoice([]string{"Test1", "Test2"}, true)
		switch answer {
		case 0:
			common.DisplayNarration("Vous avez choisi le test numéro 1.")
			common.StopAllSounds()
		case 1:
			common.DisplayDialogue("chat", "le chat", "Meow meow meow. Le choix numéro 2 !")
		case 2:
			common.DisplayNarration("Vous accédez à l'inventaire.")
		}
	}
}
