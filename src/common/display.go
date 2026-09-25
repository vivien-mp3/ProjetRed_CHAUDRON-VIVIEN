package common

import (
	"fmt"
	"time"
)

// Ici sont placé les icones des personnages qui parle
var icons = map[string]string{
	"?": "�",
	"chat": "ᓚᘏᗢ", 
	"joueur": "𐀪",
	"robot": "𖠌",
	"angel": "˚₊‧꒰ა 𓂋 ໒꒱ ‧₊˚",
}


// Ici est la fonction pour les dialogues, icon = personnage qui parle, t = le nom du personnage qui parle, m = le message
func DisplayDialogue(icon string, t string, m string) {
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]")
	fmt.Printf("|[ %s ]| -- %s o--> \n\t<< %s >>\n", icons[icon], t, m)
	time.Sleep(1 * time.Second)
}

// Ici est la fonction pour la narration, m = le message
func DisplayNarration(m string){
	fmt.Println("[{(--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--__--)}]")
	fmt.Printf("\t[[ %s ]]\n", m)
	time.Sleep(2 * time.Second)
}
// ici c'est les stat du joueur
func DisplayInfo(name string, t string, pvmax int, pv int, atk int, def int, Monnaie int) {
	fmt.Println("[{(______________________________________________________________________________________)}]")
	fmt.Printf("\t NOM DU JOUEUR: %s\n\tTYPE DU JOUEUR: %s\n", name, t)
	fmt.Printf("[{(=======================================)}]\n")
	fmt.Printf("\t POINTS DE VIES: %d / %d\n", pv, pvmax)
	fmt.Printf("[{(=======================================)}]\n")
	fmt.Printf("\t POINTS D'ATTAQUE: %d\t POINT DE DEFENSE: %d\n", atk, def)
	fmt.Printf("[{(=======================================)}]\n")
	fmt.Printf("\t MONNAIE: %d\t \n", Monnaie)
}

func DisplayBattle(tour int, ename string, epv int, epvmax int, pname string, ppv int, ppvmax int, ppr int, pprmax int) {
	fmt.Printf("\n[{(======================)}][ TOUR %d ][{(======================)}]\\nn", tour)
	fmt.Printf("\t%s | POINTS DE VIE: %d / %d |\n", pname, ppv, ppvmax)
	fmt.Printf("\t| POINTS DE RÊVES: %d / %d |\n", ppr, pprmax)
	fmt.Printf("[{(=============================================================)}]\n\n")
	fmt.Printf("\t%s | POINTS DE VIE: %d / %d |\n", ename, epv, epvmax)
	fmt.Printf("\t%s se bat encore, continue de te battre.\n", ename)
	fmt.Printf("[{(=============================================================)}]\n")
}

func DisplayTitle() {

	for result := true; result; {
		fmt.Println("\033[1;35m")
		fmt.Println(`
	 ____             _      _     ____    ___                
	|  _ \ _ __ ___  (_) ___| |_  |  _ \  /_/_\   _____  
	| |_) | '__/ _ \ | |/ _ \ __| | |_) |/ _ \ \ / / _ \  
	|  __/| | | (_) || |  __/ |_  |  _ <|  __/\ V /  __/  
	|_|   |_|  \___/_/ |\___|\__| |_| \_\\___| \_/ \___|  
	               |__/                                    
	`)
	}
}
