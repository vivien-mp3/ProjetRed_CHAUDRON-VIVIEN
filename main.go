package main

import (
	"projet/src"
)

type character struct {
	pc int
}

func main() {

	src.DisplayDialogue("chat", "Le chat", "Je suis le meow")
	src.DisplayDialogue("joueur", "Le joueur", "Oe oe oe !")
	mess := "Un vent froid coule sur votre peau.\n\tVous vous sentez bizarre, le front chaud, mais le corps froid.\n\tCela semble comme un cauchemard."
	src.DisplayNarration(mess)
}
