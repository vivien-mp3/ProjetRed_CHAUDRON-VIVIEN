package main

import (
	//"fmt"
	"projet/src"

	//"projet/src/battle"
	"projet/src/common"
	"projet/src/story"

	//"projet/src/economie"

	"projet/src/economie"
	"projet/src/player"
	"time"
)

type character struct {
	pc int
}

func main() {
	//permet initialiser le menu
	src.Startmenu()
	//requete pour le personnage
	common.DisplayDialogue("robot", "I.A Vaisseau", "Alerte! Alerte! Pilote non reconnu. Badge valide mais infos indescriptible. Veuillez vous identifier.")
	plrName := player.EnterName()
	plrType := player.EnterType()

	//information sur le personnage
	var plr player.Character
	plr.InitCharacter(plrName, plrType)

	// permet de lancer la narration
	common.DisplayInfo(plr.NAME, plr.TYPE, plr.PVMAX, plr.PV, plr.ATK, plr.DEF, plr.MONNAIE)

	time.Sleep(5 * time.Second)

	common.DisplayDialogue("robot", "I.A Vaisseau", "Alerte! Alerte! Crash imminent, préparez vous à la collision dans... 3.. 2.. 1.")
	
	time.Sleep(5 * time.Second)

	story.BeginAdventure(&plr)

	time.Sleep(1 * time.Second)
	common.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	time.Sleep(2 * time.Second)
	common.DisplayDialogue("joueur", player.NomSaisi, "Oestory. oe oe !")
	time.Sleep(1 * time.Second)
	mess := "Un vent froid coule sur votre peau.\n\tVous vous sentez bizarre, le front chaud, mais le corps froid.\n\tCela semble comme un cauchemard."
	common.DisplayNarration(mess)
	src.Carte(&plr)
	// Rencontre avec le Marchand
	economie.Boutique(&plr)
	player.AccesInventory(&plr.INV)
}
