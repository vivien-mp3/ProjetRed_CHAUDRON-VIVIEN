package main

import (
	"bufio"
	"fmt"
	"os"
	"projet/src"

	"projet/src/common"
	"strings"
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
	src.StartMenu()
	//choix du type de personnage
	fmt.Println("Veuillez Saisir le type de personnage entre les différents proposé :")
	fmt.Println("1 - Humain || 2 - chartaceus || 3 - plantyrien")
	var typesaisi string
	firstvalue := bufio.NewScanner(os.Stdin)
	firstvalue.Scan()
	typesaisi = strings.TrimSpace(firstvalue.Text())
	//choix du nom du personnage
	fmt.Println("Veuillez Saisir le nom de votre personnage entre les différents proposé :")
	var nomsaisi string
	secondvalue := bufio.NewScanner(os.Stdin)
	secondvalue.Scan()
	nomsaisi = strings.TrimSpace(secondvalue.Text())
	//information sur le personnage
	var test src.Character
	test.InitCharacter(nomsaisi, typesaisi)
	fmt.Println(test)
	// permet de lancer la narration
	src.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	src.DisplayDialogue("joueur", "Le joueur", "Oe oe oe !")
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
