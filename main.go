package main

import (
	"bufio"
	"fmt"
	"os"
	"projet/src"
	"strings"
)

type character struct {
	pc int
}

func main() {
	//permet initialiser le menu
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
	src.DisplayNarration(mess)
	src.Carte()
}
