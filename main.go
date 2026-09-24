package main

import (
	//"fmt"
	"projet/src"
	"projet/src/battle"
	"projet/src/common"
	//"projet/src/economie"
	"projet/src/player"
	chrfol "projet/src/player"
	"time"
)

type character struct {
	pc int
}

func main() {
	//permet initialiser le menu
	src.Startmenu()
	//requete pour le personnage
	player.AddInventory("épée")
	player.AddInventory("couteau")
	player.AddInventory("épée")
	player.AccesInventory()
	plrName := chrfol.EnterName()
	plrType := chrfol.EnterType()

	//information sur le personnage
	var plr chrfol.Character
	plr.InitCharacter(plrName, plrType)

	var dummy battle.AI
	battle.StartBattle(&plr, dummy, "dummy", 100, 10)

	// permet de lancer la narration
	common.DisplayInfo(plr.NAME, plr.TYPE, plr.PVMAX, plr.PV, plr.ATK, plr.DEF, plr.MONNAIE)
	time.Sleep(1 * time.Second)
	common.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	time.Sleep(2 * time.Second)
	common.DisplayDialogue("joueur", chrfol.NomSaisi, "Oestory. oe oe !")
	time.Sleep(1 * time.Second)
	mess := "Un vent froid coule sur votre peau.\n\tVous vous sentez bizarre, le front chaud, mais le corps froid.\n\tCela semble comme un cauchemard."
	common.DisplayNarration(mess)
	src.Carte(&plr)
	// Rencontre avec le Marchand
	//economie.Boutique(&plr)
	//player.AccesInventory()
}
